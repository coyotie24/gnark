/*
Copyright © 2020 ConsenSys

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sw

import (
	"testing"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/gadgets/algebra/fields"
	"github.com/consensys/gurvy"
	"github.com/consensys/gurvy/bls377"
)

type lineEvalBLS377 struct {
	Q, R G2Jac
	P    G1Jac `gnark:",public"`
}

func (circuit *lineEvalBLS377) Define(curveID gurvy.ID, cs *frontend.CS) error {
	var expected LineEvalRes
	LineEvalBLS377(cs, circuit.Q, circuit.R, circuit.P, &expected, fields.GetBLS377ExtensionFp12(cs))
	cs.MUSTBE_EQ(expected.r0.A0, "220291599185938038585565774521033812062947190299680306664648725201730830885666933651848261361463591330567860207241")
	cs.MUSTBE_EQ(expected.r0.A1, "232134458700276476669584229661634543747068594368664068937164975724095736595288995356706959089579876199020312643174")
	cs.MUSTBE_EQ(expected.r1.A0, "74241662856820718491669277383162555524896537826488558937227282983357670568906847284642533051528779250776935382660")
	cs.MUSTBE_EQ(expected.r1.A1, "9787836945036920457066634104342154603142239983688979247440278426242314457905122599227144555989168817796094251258")
	cs.MUSTBE_EQ(expected.r2.A0, "85129589817387660717039592198118788807152207633847410148299763250229022303850156734979397272700502238285752744807")
	cs.MUSTBE_EQ(expected.r2.A1, "245761211327131018855579902758747359135620549826797077633679496719449586668701082009536667506317412690997533857875")

	return nil
}

func TestLineEvalBLS377(t *testing.T) {

	// create the cs
	var circuit, witness lineEvalBLS377
	r1cs, err := frontend.Compile(gurvy.BW761, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	var Q, R bls377.G2Jac
	var P bls377.G1Jac

	Q.X.A0.SetString("11467063222684898633036104763692544506257812867640109164430855414494851760297509943081481005947955008078272733624")
	Q.X.A1.SetString("153924906120314059329163510034379429156688480181182668999642334674073859906019623717844462092443710331558842221198")
	Q.Y.A0.SetString("217426664443013466493849511677243421913435679616098405782168799962712362374085608530270502677771125796970144049342")
	Q.Y.A1.SetString("220113305559851867470055261956775835250492241909876276448085325823827669499391027597256026508256704101389743638320")
	Q.Z.A0.SetOne()

	R.X.A0.SetString("38348804106969641131654336618231918247608720362924380120333996440589719997236048709530218561145001033408367199467")
	R.X.A1.SetString("208837221672103828632878568310047865523715993428626260492233587961023171407529159232705047544612759994485307437530")
	R.Y.A0.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	R.Y.A1.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")
	R.Z.A0.SetOne()

	P.X.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	P.Y.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")
	P.Z.SetOne()

	witness.Q.Assign(&Q)
	witness.R.Assign(&R)
	witness.P.Assign(&P)

	assert := groth16.NewAssert(t)
	assert.CorrectExecution(r1cs, &witness, nil)
}

type lineEvalAffineBLS377 struct {
	Q, R G2Affine
	P    G1Affine `gnark:",public"`
}

func (circuit *lineEvalAffineBLS377) Define(curveID gurvy.ID, cs *frontend.CS) error {
	var expected LineEvalRes
	LineEvalAffineBLS377(cs, circuit.Q, circuit.R, circuit.P, &expected, fields.GetBLS377ExtensionFp12(cs))
	cs.MUSTBE_EQ(expected.r0.A0, "220291599185938038585565774521033812062947190299680306664648725201730830885666933651848261361463591330567860207241")
	cs.MUSTBE_EQ(expected.r0.A1, "232134458700276476669584229661634543747068594368664068937164975724095736595288995356706959089579876199020312643174")
	cs.MUSTBE_EQ(expected.r1.A0, "74241662856820718491669277383162555524896537826488558937227282983357670568906847284642533051528779250776935382660")
	cs.MUSTBE_EQ(expected.r1.A1, "9787836945036920457066634104342154603142239983688979247440278426242314457905122599227144555989168817796094251258")
	cs.MUSTBE_EQ(expected.r2.A0, "85129589817387660717039592198118788807152207633847410148299763250229022303850156734979397272700502238285752744807")
	cs.MUSTBE_EQ(expected.r2.A1, "245761211327131018855579902758747359135620549826797077633679496719449586668701082009536667506317412690997533857875")

	return nil
}

func TestLineEvalAffineBLS377(t *testing.T) {

	// create the cs
	var circuit, witness lineEvalAffineBLS377
	r1cs, err := frontend.Compile(gurvy.BW761, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	var Q, R bls377.G2Affine
	var P bls377.G1Affine

	Q.X.A0.SetString("11467063222684898633036104763692544506257812867640109164430855414494851760297509943081481005947955008078272733624")
	Q.X.A1.SetString("153924906120314059329163510034379429156688480181182668999642334674073859906019623717844462092443710331558842221198")
	Q.Y.A0.SetString("217426664443013466493849511677243421913435679616098405782168799962712362374085608530270502677771125796970144049342")
	Q.Y.A1.SetString("220113305559851867470055261956775835250492241909876276448085325823827669499391027597256026508256704101389743638320")

	R.X.A0.SetString("38348804106969641131654336618231918247608720362924380120333996440589719997236048709530218561145001033408367199467")
	R.X.A1.SetString("208837221672103828632878568310047865523715993428626260492233587961023171407529159232705047544612759994485307437530")
	R.Y.A0.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	R.Y.A1.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")

	P.X.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	P.Y.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")

	witness.Q.Assign(&Q)
	witness.R.Assign(&R)
	witness.P.Assign(&P)

	assert := groth16.NewAssert(t)
	assert.CorrectExecution(r1cs, &witness, nil)
}

type pairingAffineBLS377 struct {
	Q          G2Affine
	P          G1Affine `gnark:",public"`
	pairingRes bls377.PairingResult
}

func (circuit *pairingAffineBLS377) Define(curveID gurvy.ID, cs *frontend.CS) error {

	ateLoop := uint64(9586122913090633729)
	ext := fields.GetBLS377ExtensionFp12(cs)
	pairingInfo := PairingContext{AteLoop: ateLoop, Extension: ext}

	milRes := fields.E12{}
	pairingRes := fields.E12{}

	MillerLoopAffine(cs, circuit.P, circuit.Q, &milRes, pairingInfo)
	pairingRes.FinalExpoBLS(cs, &milRes, ateLoop, ext)

	mustbeEq(cs, pairingRes, &circuit.pairingRes, "pairingres")

	return nil
}

func TestPairingAffineBLS377(t *testing.T) {
	P, Q, pairingRes := pairingData()

	// create cs
	var circuit, witness pairingAffineBLS377
	circuit.pairingRes = pairingRes
	r1cs, err := frontend.Compile(gurvy.BW761, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	// set the cs
	witness.P.Assign(&P)
	witness.Q.Assign(&Q)

	assert := groth16.NewAssert(t)
	expectedValues := make(map[string]interface{})
	addExpectedFP12(&pairingRes, "pairingres", expectedValues)
	assert.CorrectExecution(r1cs, &witness, expectedValues)

}

type pairingBLS377 struct {
	Q          G2Jac
	P          G1Jac `gnark:",public"`
	pairingRes bls377.PairingResult
}

func (circuit *pairingBLS377) Define(curveID gurvy.ID, cs *frontend.CS) error {

	ateLoop := uint64(9586122913090633729)
	ext := fields.GetBLS377ExtensionFp12(cs)
	pairingInfo := PairingContext{AteLoop: ateLoop, Extension: ext}

	milRes := fields.E12{}
	MillerLoop(cs, circuit.P, circuit.Q, &milRes, pairingInfo)

	pairingRes := fields.E12{}
	pairingRes.FinalExpoBLS(cs, &milRes, ateLoop, ext)

	mustbeEq(cs, pairingRes, &circuit.pairingRes, "pairingres")

	return nil
}

func TestPairingBLS377(t *testing.T) {
	// pairing test data
	_P, _Q, pairingRes := pairingData()
	var Q bls377.G2Jac
	var P bls377.G1Jac
	P.FromAffine(&_P)
	Q.FromAffine(&_Q)

	// create cs
	var circuit, witness pairingBLS377
	circuit.pairingRes = pairingRes
	r1cs, err := frontend.Compile(gurvy.BW761, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	// assign values to witness
	witness.P.Assign(&P)
	witness.Q.Assign(&Q)

	assert := groth16.NewAssert(t)
	expectedValues := make(map[string]interface{})
	addExpectedFP12(&pairingRes, "pairingres", expectedValues)
	assert.CorrectExecution(r1cs, &witness, expectedValues)

}

func pairingData() (P bls377.G1Affine, Q bls377.G2Affine, pairingRes bls377.PairingResult) {
	P.X.SetString("68333130937826953018162399284085925021577172705782285525244777453303237942212457240213897533859360921141590695983")
	P.Y.SetString("243386584320553125968203959498080829207604143167922579970841210259134422887279629198736754149500839244552761526603")

	Q.X.A0.SetString("129200027147742761118726589615458929865665635908074731940673005072449785691019374448547048953080140429883331266310")
	Q.X.A1.SetString("218164455698855406745723400799886985937129266327098023241324696183914328661520330195732120783615155502387891913936")
	Q.Y.A0.SetString("178797786102020318006939402153521323286173305074858025240458924050651930669327663166574060567346617543016897467207")
	Q.Y.A1.SetString("246194676937700783734853490842104812127151341609821057456393698060154678349106147660301543343243364716364400889778")

	milRes := bls377.MillerLoop(P, Q)
	pairingRes = bls377.FinalExponentiation(milRes)

	return
}

func addExpectedFP12(e12 *bls377.PairingResult, tagPrefix string, expectedValues map[string]interface{}) {
	expectedValues[tagPrefix+".C0.B0.A0"] = e12.C0.B0.A0
	expectedValues[tagPrefix+".C0.B0.A1"] = e12.C0.B0.A1
	expectedValues[tagPrefix+".C0.B1.A0"] = e12.C0.B1.A0
	expectedValues[tagPrefix+".C0.B1.A1"] = e12.C0.B1.A1
	expectedValues[tagPrefix+".C0.B2.A0"] = e12.C0.B2.A0
	expectedValues[tagPrefix+".C0.B2.A1"] = e12.C0.B2.A1
	expectedValues[tagPrefix+".C1.B0.A0"] = e12.C1.B0.A0
	expectedValues[tagPrefix+".C1.B0.A1"] = e12.C1.B0.A1
	expectedValues[tagPrefix+".C1.B1.A0"] = e12.C1.B1.A0
	expectedValues[tagPrefix+".C1.B1.A1"] = e12.C1.B1.A1
	expectedValues[tagPrefix+".C1.B2.A0"] = e12.C1.B2.A0
	expectedValues[tagPrefix+".C1.B2.A1"] = e12.C1.B2.A1
}

func mustbeEq(cs *frontend.CS, fp12 fields.E12, e12 *bls377.PairingResult, tagPrefix string) {
	cs.Tag(fp12.C0.B0.A0, tagPrefix+".C0.B0.A0")
	cs.Tag(fp12.C0.B0.A1, tagPrefix+".C0.B0.A1")
	cs.Tag(fp12.C0.B1.A0, tagPrefix+".C0.B1.A0")
	cs.Tag(fp12.C0.B1.A1, tagPrefix+".C0.B1.A1")
	cs.Tag(fp12.C0.B2.A0, tagPrefix+".C0.B2.A0")
	cs.Tag(fp12.C0.B2.A1, tagPrefix+".C0.B2.A1")
	cs.Tag(fp12.C1.B0.A0, tagPrefix+".C1.B0.A0")
	cs.Tag(fp12.C1.B0.A1, tagPrefix+".C1.B0.A1")
	cs.Tag(fp12.C1.B1.A0, tagPrefix+".C1.B1.A0")
	cs.Tag(fp12.C1.B1.A1, tagPrefix+".C1.B1.A1")
	cs.Tag(fp12.C1.B2.A0, tagPrefix+".C1.B2.A0")
	cs.Tag(fp12.C1.B2.A1, tagPrefix+".C1.B2.A1")
	cs.MUSTBE_EQ(fp12.C0.B0.A0, e12.C0.B0.A0)
	cs.MUSTBE_EQ(fp12.C0.B0.A1, e12.C0.B0.A1)
	cs.MUSTBE_EQ(fp12.C0.B1.A0, e12.C0.B1.A0)
	cs.MUSTBE_EQ(fp12.C0.B1.A1, e12.C0.B1.A1)
	cs.MUSTBE_EQ(fp12.C0.B2.A0, e12.C0.B2.A0)
	cs.MUSTBE_EQ(fp12.C0.B2.A1, e12.C0.B2.A1)
	cs.MUSTBE_EQ(fp12.C1.B0.A0, e12.C1.B0.A0)
	cs.MUSTBE_EQ(fp12.C1.B0.A1, e12.C1.B0.A1)
	cs.MUSTBE_EQ(fp12.C1.B1.A0, e12.C1.B1.A0)
	cs.MUSTBE_EQ(fp12.C1.B1.A1, e12.C1.B1.A1)
	cs.MUSTBE_EQ(fp12.C1.B2.A0, e12.C1.B2.A0)
	cs.MUSTBE_EQ(fp12.C1.B2.A1, e12.C1.B2.A1)
}
