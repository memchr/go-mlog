package impl

import (
	"github.com/Vilsol/go-mlog/m"
	"github.com/Vilsol/go-mlog/transpiler"
	"strings"
)

func init() {
	transpiler.RegisterSelector("m.RTAny", m.RTAny)
	transpiler.RegisterSelector("m.RTEnemy", m.RTEnemy)
	transpiler.RegisterSelector("m.RTAlly", m.RTAlly)
	transpiler.RegisterSelector("m.RTPlayer", m.RTPlayer)
	transpiler.RegisterSelector("m.RTAttacker", m.RTAttacker)
	transpiler.RegisterSelector("m.RTFlying", m.RTFlying)
	transpiler.RegisterSelector("m.RTBoss", m.RTBoss)
	transpiler.RegisterSelector("m.RTGround", m.RTGround)

	transpiler.RegisterSelector("m.RSDistance", m.RSDistance)
	transpiler.RegisterSelector("m.RSHealth", m.RSHealth)
	transpiler.RegisterSelector("m.RSShield", m.RSShield)
	transpiler.RegisterSelector("m.RSArmor", m.RSArmor)
	transpiler.RegisterSelector("m.RSMaxHealth", m.RSMaxHealth)

	transpiler.RegisterSelector("m.BCore", m.BCore)
	transpiler.RegisterSelector("m.BStorage", m.BStorage)
	transpiler.RegisterSelector("m.BGenerator", m.BGenerator)
	transpiler.RegisterSelector("m.BTurret", m.BTurret)
	transpiler.RegisterSelector("m.BFactory", m.BFactory)
	transpiler.RegisterSelector("m.BRepair", m.BRepair)
	transpiler.RegisterSelector("m.BRally", m.BRally)
	transpiler.RegisterSelector("m.BBattery", m.BBattery)
	transpiler.RegisterSelector("m.BResupply", m.BResupply)
	transpiler.RegisterSelector("m.BReactor", m.BReactor)
	transpiler.RegisterSelector("m.BUnitModifier", m.BUnitModifier)
	transpiler.RegisterSelector("m.BExtinguisher", m.BExtinguisher)

	transpiler.RegisterSelector("m.This", "@this")
	transpiler.RegisterSelector("m.ThisX", "@thisx")
	transpiler.RegisterSelector("m.ThisXf", "@thisx")
	transpiler.RegisterSelector("m.ThisY", "@thisy")
	transpiler.RegisterSelector("m.ThisYf", "@thisy")
	transpiler.RegisterSelector("m.Ipt", "@ipt")
	transpiler.RegisterSelector("m.Counter", "@counter")
	transpiler.RegisterSelector("m.Links", "@links")
	transpiler.RegisterSelector("m.CurUnit", "@unit")
	transpiler.RegisterSelector("m.Time", "@time")
	transpiler.RegisterSelector("m.Tick", "@tick")
	transpiler.RegisterSelector("m.MapW", "@mapw")
	transpiler.RegisterSelector("m.MapH", "@maph")

	// HealthC's attributes
	transpiler.RegisterFuncTranslation("Health", createSensorFuncTranslation("@health"))
	transpiler.RegisterFuncTranslation("Name", createSensorFuncTranslation("@name"))
	transpiler.RegisterFuncTranslation("X", createSensorFuncTranslation("@x"))
	transpiler.RegisterFuncTranslation("Y", createSensorFuncTranslation("@y"))

	transpiler.RegisterFuncTranslation("TotalItems", createSensorFuncTranslation("@totalItems"))
	transpiler.RegisterFuncTranslation("ItemCapacity", createSensorFuncTranslation("@itemCapacity"))
	transpiler.RegisterFuncTranslation("Rotation", createSensorFuncTranslation("@rotation"))
	transpiler.RegisterFuncTranslation("ShootX", createSensorFuncTranslation("@shootX"))
	transpiler.RegisterFuncTranslation("ShootY", createSensorFuncTranslation("@shootY"))
	transpiler.RegisterFuncTranslation("Shooting", createSensorFuncTranslation("@shooting"))

	// Building's attributes
	transpiler.RegisterFuncTranslation("TotalLiquids", createSensorFuncTranslation("@totalLiquids"))
	transpiler.RegisterFuncTranslation("LiquidCapaticy", createSensorFuncTranslation("@liquidCapaticy"))
	transpiler.RegisterFuncTranslation("TotalPower", createSensorFuncTranslation("@totalPower"))
	transpiler.RegisterFuncTranslation("PowerCapaticy", createSensorFuncTranslation("@powerCapaticy"))
	transpiler.RegisterFuncTranslation("PowerNetStored", createSensorFuncTranslation("@powerNetStored"))
	transpiler.RegisterFuncTranslation("PowerNetCapacity", createSensorFuncTranslation("@powerNetCapacity"))
	transpiler.RegisterFuncTranslation("PowerNetIn", createSensorFuncTranslation("@powerNetIn"))
	transpiler.RegisterFuncTranslation("PowerNetOut", createSensorFuncTranslation("@powerNetOut"))
	transpiler.RegisterFuncTranslation("Heat", createSensorFuncTranslation("@heat"))
	transpiler.RegisterFuncTranslation("Efficiency", createSensorFuncTranslation("@efficiency"))
	transpiler.RegisterFuncTranslation("Enabled", createSensorFuncTranslation("@enabled"))
}

func createSensorFuncTranslation(attribute string) transpiler.Translator {
	return transpiler.Translator{
		Count: func(args []transpiler.Resolvable, vars []transpiler.Resolvable) int {
			return 1
		},
		Variables: 1,
		Translate: func(args []transpiler.Resolvable, vars []transpiler.Resolvable) ([]transpiler.MLOGStatement, error) {
			return []transpiler.MLOGStatement{
				&transpiler.MLOG{
					Statement: [][]transpiler.Resolvable{
						{
							&transpiler.Value{Value: "sensor"},
							vars[0],
							&transpiler.Value{Value: strings.Trim(args[0].GetValue(), "\"")},
							&transpiler.Value{Value: attribute},
						},
					},
				},
			}, nil
		},
	}
}

func genBasicFuncTranslation(constants []string, nArgs int, nVars int) transpiler.TranslateFunc {
	return func(args []transpiler.Resolvable, vars []transpiler.Resolvable) ([]transpiler.MLOGStatement, error) {
		statements := make([]transpiler.Resolvable, len(constants)+nArgs+nVars)

		for i, constant := range constants {
			statements[i] = &transpiler.Value{Value: constant}
		}

		for i := 0; i < nArgs; i++ {
			statements[i+len(constants)] = &transpiler.Value{Value: args[i].GetValue()}
		}

		for i := 0; i < nVars; i++ {
			statements[i+len(constants)+nArgs] = vars[i]
		}

		return []transpiler.MLOGStatement{
			&transpiler.MLOG{
				Statement: [][]transpiler.Resolvable{statements},
			},
		}, nil
	}
}

const (
	TRUE  = "true"
	FALSE = "false"
)
