package main

import "github.com/ArchRobison/FrequonInvaders/phrase"

// grammar describes the grammar for teletype technobabble.
var grammar = phrase.MakeGrammar([]string{
	"1: Boot %A %O",
	"1: Reset %A %O",
	"1: Load %A %O",
	"2: Enable %A %O",
	"2: Engage %A %O",
	"2: Digitize %A %O",
	"3: Toggle %A %O",
	"3: Activate %A %O",
	"4: Invert %A %O",
	"4: Turn on %A %O",
	"4: Spin up %A %O",
	"5: Tune %A %O",
	"5: Rev up %A %O",
	"6: Switch on %A %O",
	"6: Check %A %O",
	"6: Lock %A %O",
	"7: Power on %A %O",
	"7: Energize %A %O",
	"8: Compress %A %O",
	"O: overthruster",
	"O: inverter",
	"O: integrator",
	"O: differentiator",
	"O: accumulator",
	"O: accelerator",
	"O: processors",
	"O: twistor",
	"O: eigenvector",
	"O: function",
	"O: reactor",
	"O: transporter",
	"A: micro",
	"A: elliptic",
	"A: interprocedural",
	"A: hyperfinite",
	"A: photonic",
	"A: quantum",
	"A: nano",
	"A: alpha",
	"A: beta",
	"A: psi",
	"A: transwarp",
	"A: harmonic",
	"A: complex",
	"A: hyperbolic",
	"A: pseudomorphic",
	"A: acoustic",
	"S: Score = $",
	"H: Your score of $ merits recording.",
	"W: You Win!",
})
