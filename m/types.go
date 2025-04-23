package m

type RadarTarget = string

const (
	// Target anything
	RTAny      = RadarTarget("any")
	RTEnemy    = RadarTarget("enemy")
	RTAlly     = RadarTarget("ally")
	RTPlayer   = RadarTarget("player")
	RTAttacker = RadarTarget("attacker")
	RTFlying   = RadarTarget("flying")
	RTBoss     = RadarTarget("boss")
	RTGround   = RadarTarget("ground")
)

type RadarSort = string

const (
	RSDistance  = RadarSort("distance")
	RSHealth    = RadarSort("health")
	RSShield    = RadarSort("shield")
	RSArmor     = RadarSort("armor")
	RSMaxHealth = RadarSort("maxHealth")
)

type Link = interface{}

type HealthC = interface {
	// Object's content name.
	Type() string

	// the current health value of Unit or Building
	Health() int

	// the name of Unit or Building
	Name() string

	// the X coordinate of Unit or Building
	X() float64

	// the Y coordinate of Unit or Building
	Y() float64

	//  the total number of items stored in Unit or Building
	TotalItems() int

	//  the capacity of items stored in Unit or Building
	ItemCapacity() int

	//  the angle of view of Unit or Building
	Rotation() float64

	//  the X coordinate of Unit or Building's collimator
	ShootX() float64

	//  the Y coordinate of Unit or Building's collimator
	ShootY() float64

	// Check whether Unit or Building fires
	Shooting() bool
}

type Ranged = interface{}

type Unit = interface {
	HealthC
	Ranged
}

type Building = interface {
	HealthC
	Ranged
	Items
	Liquids

	// Total amount of liquids currently stored in the container(Building)
	TotalLiquids() float64

	// The maximum amount of liquids stored in the container(Building)
	LiquidCapaticy() float64

	// In case of unbuffered consumers, this is the percentage (1.0f = 100%) of the demanded power which can be supplied.
	// Blocks will work at a reduced efficiency if this is not equal to 1.0f.
	// In case of buffered consumers, this is storage capacity.
	TotalPower() float64

	// In case of unbuffered consumers, this is the 0
	// In case of buffered consumers, this is the maximum storage capacity.
	PowerCapaticy() float64

	// The total power currently stored in the grid (Only machines connected to the grid)
	PowerNetStored() float64

	// The maximum power capacity in the grid (Only machines connected to the grid)
	PowerNetCapacity() float64

	// The input power of the current grid (Only machines connected to the grid)
	PowerNetIn() float64

	// The output power of the current grid (Only machines connected to the grid)
	PowerNetOut() float64

	// The heat from the machine (Just Thorium Reactor)
	Heat() float64

	// Machine efficiency
	Efficiency() float64

	// Check whether the machine is available
	Enabled() bool
}

var (
	// A Building Object that represents the processor itself.
	// You can use this with sensor to find various properties about the processor.
	This = Building(nil)

	// The x coordinate of the processor.
	ThisX = 0

	// The y coordinate of the processor.
	ThisY = 0

	// The number of instructions executed per tick (60 ticks/second).
	//
	// Micro Processor -> 2
	// Logic Processor -> 8
	// Hyper Processor -> 25
	Ipt = 0

	// A variable that represents the next line the processor will read code from, equivalent to %IP in x86.
	// It can be changed like any other variable as another way to perform jumps.
	Counter = 0

	// A constant that equals the number of buildings linked to the processor.
	// It is changed by the processor when blocks are linked or unlinked.
	Links = 0

	// A constant that represents the current bound unit.
	// It only changes when the processor unbinds a unit, or binds another one.
	CurUnit = Unit(nil)

	// Represents the current UNIX timestamp in milliseconds.
	Time = 0

	// Represents the amount of ticks (60 ticks/second) since the map began.
	Tick = float64(0)

	// Width of the map, in tiles.
	MapW = 0

	// Height of the map, in tiles.
	MapH = 0
)

type BlockFlag = string

const (
	BCore         = BlockFlag("core")
	BStorage      = BlockFlag("storage")
	BGenerator    = BlockFlag("generator")
	BTurret       = BlockFlag("turret")
	BFactory      = BlockFlag("factory")
	BRepair       = BlockFlag("repair")
	BRally        = BlockFlag("rally")
	BBattery      = BlockFlag("battery")
	BResupply     = BlockFlag("resupply")
	BReactor      = BlockFlag("reactor")
	BUnitModifier = BlockFlag("unitModifier")
	BExtinguisher = BlockFlag("extinguisher")
)
