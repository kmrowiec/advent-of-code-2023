package day20

import (
	"fmt"
	"sort"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day20/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day20/input/input1.txt")
}

type Pulse struct {
	Source string
	Target string
	Value  string
}

type Device interface {
	GetId() string
	GetTargets() []string
	Activate(pulse Pulse) []Pulse
	GetState() string
}

type Broadcast struct {
	Id      string
	Targets []string
}

func (b *Broadcast) Activate(pulse Pulse) []Pulse {
	result := make([]Pulse, 0)
	for _, target := range b.Targets {
		result = append(result, Pulse{Target: target, Value: pulse.Value, Source: b.Id})
	}
	return result
}

func (b *Broadcast) GetState() string {
	return b.Id + ": OK"
}

func (b *Broadcast) GetId() string {
	return b.Id
}

func (b *Broadcast) GetTargets() []string {
	return b.Targets
}

type FlipFlop struct {
	Id      string
	Targets []string
	Enabled bool
}

func (f *FlipFlop) GetTargets() []string {
	return f.Targets
}

func (f *FlipFlop) Activate(pulse Pulse) []Pulse {
	result := make([]Pulse, 0)
	if pulse.Value == "L" {
		f.Enabled = !f.Enabled
		value := "L"
		if f.Enabled {
			value = "H"
		}
		for _, target := range f.Targets {
			result = append(result, Pulse{Target: target, Value: value, Source: f.Id})
		}
	}
	return result
}

func (f *FlipFlop) GetState() string {
	return fmt.Sprintf("%v(%v)", f.Id, f.Enabled)
}

func (f *FlipFlop) GetId() string {
	return f.Id
}

type Conjunction struct {
	Id      string
	Targets []string
	Memory  map[string]string
}

func (c *Conjunction) GetTargets() []string {
	return c.Targets
}

func (c *Conjunction) Activate(pulse Pulse) []Pulse {

	c.Memory[pulse.Source] = pulse.Value

	result := make([]Pulse, 0)

	value := "L"

	for _, v := range c.Memory {
		if v == "L" {
			value = "H"
			break
		}
	}

	for _, target := range c.Targets {
		result = append(result, Pulse{Target: target, Value: value, Source: c.Id})
	}
	return result
}

func (c *Conjunction) GetState() string {
	orderedMemory := make([]string, 0)
	for k, v := range c.Memory {
		orderedMemory = append(orderedMemory, k+":"+v)
	}
	sort.SliceStable(orderedMemory, func(i, j int) bool {
		return strings.Compare(orderedMemory[i], orderedMemory[j]) == -1
	})
	return fmt.Sprintf("%v(%v)", c.Id, orderedMemory)
}

func (c *Conjunction) GetId() string {
	return c.Id
}

func CreateDevice(line string) (string, Device) {
	namePart, targetsPart, _ := strings.Cut(line, " -> ")
	targets := strings.Split(targetsPart, ", ")
	if namePart == "broadcaster" {
		return "broadcaster", &Broadcast{Id: "broadcaster", Targets: targets}
	}
	deviceType, id := namePart[0], namePart[1:]
	switch deviceType {
	case '%':
		return id, &FlipFlop{Id: id, Targets: targets}
	case '&':
		return id, &Conjunction{Id: id, Targets: targets, Memory: map[string]string{}}
	}
	panic("Unknown device!")
}

func GetMachineState(machine map[string]Device) string {
	result := ""
	orderedDevices := make([]Device, 0)
	for _, device := range machine {
		orderedDevices = append(orderedDevices, device)
	}
	sort.SliceStable(orderedDevices, func(i, j int) bool {
		return strings.Compare(orderedDevices[i].GetId(), orderedDevices[j].GetId()) == -1
	})
	for _, device := range orderedDevices {
		result += device.GetState()
	}
	return result
}

func SolvePart1(inputFile string) string {
	result := 0

	machine := map[string]Device{"output": &Broadcast{Id: "output"}}

	// Create all devices
	for _, line := range helper.ReadInputFile(inputFile) {
		id, d := CreateDevice(line)
		machine[id] = d
	}

	// Initialise conjunction devices - find all their inputs and set the map
	for id, device := range machine {
		for _, targetId := range device.GetTargets() {
			if c, ok := machine[targetId].(*Conjunction); ok {
				c.Memory[id] = "L"
			}
		}
	}

	pulseQueue := make([]Pulse, 0)
	pulseCounts := map[string]int{"H": 0, "L": 0}

	fmt.Println(machine)

	initialState := GetMachineState(machine)
	fmt.Printf("Initial state: %v\n", initialState)

	i := 1
	for ; i <= 1000; i++ {
		// Initial pulse from the button
		pulseQueue = append(pulseQueue, Pulse{Source: "button", Target: "broadcaster", Value: "L"})
		pulseCounts["L"] = pulseCounts["L"] + 1

		for len(pulseQueue) > 0 {
			currentPulse := pulseQueue[0]
			pulseQueue = pulseQueue[1:]
			fmt.Printf("%v - %v -> %v\n", currentPulse.Source, currentPulse.Value, currentPulse.Target)

			if device, ok := machine[currentPulse.Target]; ok {
				newPulses := device.Activate(currentPulse)
				for _, p := range newPulses {
					pulseCounts[p.Value] = pulseCounts[p.Value] + 1
					pulseQueue = append(pulseQueue, p)
				}
			}
		}

		currentState := GetMachineState(machine)
		fmt.Printf("Current state: %v\n", currentState)

		if currentState == initialState {
			fmt.Printf("The state repeated after %d button clicks!\n", i)
			break
		}
	}

	multiplier := 1000 / i

	if i == 1001 {
		fmt.Println("Did not find repeated state!")
		multiplier = 1
	}

	result = pulseCounts["H"] * multiplier * pulseCounts["L"] * multiplier
	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {

	machine := map[string]Device{"output": &Broadcast{Id: "output"}}

	// Create all devices
	for _, line := range helper.ReadInputFile(inputFile) {
		id, d := CreateDevice(line)
		machine[id] = d
	}

	// Initialise conjunction devices - find all their inputs and set the map
	for id, device := range machine {
		for _, targetId := range device.GetTargets() {
			if c, ok := machine[targetId].(*Conjunction); ok {
				c.Memory[id] = "L"
			}
		}
	}

	pulseQueue := make([]Pulse, 0)
	pulseCounts := map[string]int{"H": 0, "L": 0}

	fmt.Println(machine)

	initialState := GetMachineState(machine)
	fmt.Printf("Initial state: %v\n", initialState)

	highInputsMap := map[string]int{}

	i := 1
	for ; len(highInputsMap) < 4; i++ {
		// Initial pulse from the button
		pulseQueue = append(pulseQueue, Pulse{Source: "button", Target: "broadcaster", Value: "L"})
		pulseCounts["L"] = pulseCounts["L"] + 1

		for len(pulseQueue) > 0 {
			currentPulse := pulseQueue[0]
			pulseQueue = pulseQueue[1:]
			fmt.Printf("%v - %v -> %v\n", currentPulse.Source, currentPulse.Value, currentPulse.Target)

			if currentPulse.Target == "pq" && currentPulse.Value == "L" {
				highInputsMap["pq"] = i
			}

			if currentPulse.Target == "fg" && currentPulse.Value == "L" {
				highInputsMap["fg"] = i
			}

			if currentPulse.Target == "dk" && currentPulse.Value == "L" {
				highInputsMap["dk"] = i
			}

			if currentPulse.Target == "fm" && currentPulse.Value == "L" {
				highInputsMap["fm"] = i
			}

			if device, ok := machine[currentPulse.Target]; ok {
				newPulses := device.Activate(currentPulse)
				for _, p := range newPulses {
					pulseCounts[p.Value] = pulseCounts[p.Value] + 1
					pulseQueue = append(pulseQueue, p)
				}
			}
		}
	}

	return fmt.Sprint(highInputsMap)
}
