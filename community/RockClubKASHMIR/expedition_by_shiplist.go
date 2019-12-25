//---This script is created by RockClubKASHMIR---
//---The will find automatically the planet or moon contains ships from your ships list---
fromSystem = 1 // Your can change this value as you wish! If fromSystem is equal to toSystem, the script will send ships only to 1 sollar system
toSystem = 339 // Your can change this value as you wish! If fromSystem is equal to toSystem, the script will send ships only to 1 sollar system
shipsList = {PATHFINDER: 5, LARGECARGO: 300, ESPIONAGEPROBE: 11, BOMBER: 1, DESTROYER: 1}// Your can change ENTYRE List, even to left only 1 type of ships!
//=== Please, DO NOT change anything below! ===
origin = nil
master = 0
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    flts = 0
    for ShipID in shipsList {
        if ships.ByID(ShipID) > 0 {
            flts = flts + ships.ByID(ShipID)
        }
    }
    if flts > master {
        master = flts
        origin = celestial // Your Planet(or Moon)
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    for System = fromSystem; System <= toSystem; System++ {
        ships, _ = origin.GetShips()
        Destination = NewCoordinate(origin.GetCoordinate().Galaxy, System, 16, PLANET_TYPE)
        if GetSlots().ExpInUse < GetSlots().ExpTotal {
            if Destination != 0 {
                f = NewFleet()
                f.SetOrigin(origin)
                f.SetDestination(Destination)
                f.SetSpeed(HUNDRED_PERCENT)
                f.SetMission(EXPEDITION)
                for id, nbr in shipsList {
                    if ships.ByID(id) > 0 {
                        if ships.ByID(id) < nbr {nbr = ships.ByID(id)}
                        f.AddShips(id, nbr)
                    }
                }
                a, b = f.SendNow()
                if b == nil {
                    Print("The fleet is sended successfully to "+Destination)
                } else {
                    Print("The fleet is NOT sended! "+b)
                    break
                }
            } else {Print("Skipping this coordinates")}
            Sleep(Random(3000, 6000))
        } else {
            for GetSlots().ExpInUse == GetSlots().ExpTotal {
                Print("All Expedition slots are busy now! Please, wait "+ShortDur(120))
                Sleep(120000)
                GetSlots()
            }
        }
        if System >= toSystem {System = fromSystem-1}
    }
} else {Print("You don't have any ships from the desired list of ships!")}