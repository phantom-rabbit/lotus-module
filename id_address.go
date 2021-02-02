package modules

import (
	"context"
	"gorm.io/gorm/clause"
	"time"

	"gorm.io/gorm"

	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	_init "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
	cw_util "github.com/filecoin-project/lotus/cmd/lotus-chainwatch/util"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var log = logging.Logger("IdAddressTable")

type IdAddressTable struct {
	Id          string `gorm:"index;not null;"`
	Address     string `gorm:"index;not null;"`
	AddressName string

	AddressTag  string
}

func (IdAddressTable) TableName() string {
	return prefix+ "id_address"
}

func CreateIdAddressTable(db *gorm.DB) error {
	return db.AutoMigrate(&IdAddressTable{})
}

func (ida IdAddressTable) insert() error {
	err := DB.Exec(`insert into `+ida.TableName()+`(id, address, address_name) values (?, ?, ?) on conflict do nothing`, ida.Id, ida.Address, ida.AddressName).Error
	return err
}

func Sync(node api.FullNode) {
	go func() {
		for {
			now := time.Now()
			log.Debug("start sync IdAddressTable")

			initActor, err := node.StateGetActor(context.TODO(), builtin.InitActorAddr, types.EmptyTSK)
			if err != nil {
				log.Warn("node.StateGetActor failed:", err)
				return
			}

			initActorState, err := _init.Load(cw_util.NewAPIIpldStore(context.TODO(), node), initActor)
			if err != nil {
				log.Warn(err)
				return
			}

			err = initActorState.ForEachActor(func(id abi.ActorID, addr address.Address) error {
				idAddr, err := address.NewIDAddress(uint64(id))
				if err != nil {
					return err
				}

				actor, err := node.StateGetActor(context.TODO(), idAddr, types.EmptyTSK)
				name := "<undefined>"
				if err == nil {
					name = builtin.ActorNameByCode(actor.Code)
				}

				idMap := IdAddressTable{
					Id:      idAddr.String(),
					Address: addr.String(),
					AddressName: name,
					AddressTag: idAddr.String(),
				}

				err = DB.Clauses(clause.OnConflict{
					Columns: []clause.Column{{Name: "id"}},
					DoUpdates: clause.AssignmentColumns([]string{
						"address",
						"address_name",
					}),
				}).Create(&idMap).Error
				if err != nil {
					return err
				}
				return nil
			})

			if err != nil {
				log.Warn("initActorState.ForEachActor failed:", err)
				return
			}

			log.Debug("end sync IdAddressTable, time consuming:", time.Since(now).String())

			time.Sleep(30 * time.Minute)
		}
	}()
}