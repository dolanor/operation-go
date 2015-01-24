// START OMIT
/***********************
 * Mission: All Aboard *
 ***********************
 *
 * Logging into the agency network, you're quickly into the operation manifest code.
 * Sure enough, you're listed with no equipment, a typical agency snafu.
 * You're going to need full tactical gear for this mission.
 * Let's go ahead and fix that.
 *
 */

package main

import(
  "fmt"
)

func main() {
  total_agents := 5
  agents := make([]Agent,0)
  agents = append(agents, Agent{name: "Jason Marshall", equipment: "none"})
  agents = append(agents, Agent{name: "Stone Boswell", equipment: "full"})
  agents = append(agents, Agent{name: "Jane Johnson", equipment: "full"})
  agents = append(agents, Agent{name: "Max Carter", equipment: "full"})
  agents = append(agents, Agent{name: "Kay White", equipment: "full"})
// EDITABLE OMIT

  // Your code

// UNEDITABLE OMIT
  printGearTable(agents, total_agents)
}

func printGearTable(agents []Agent, total_agents int) {
  if len(agents) != total_agents {
    fmt.Println("ERROR: We can only go with", total_agents, "agents.")
    return
  }
  fmt.Println("Operation Go: Agent Manifest")
  fmt.Println("----------------------------")
  for num, agent := range agents {
    fmt.Println(num+1, agent.name, "| Gear:", agent.equipment)
  }
}

type Agent struct {
  name string
  equipment string
}
// END OMIT
