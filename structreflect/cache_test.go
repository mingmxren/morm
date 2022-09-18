package structreflect

import (
	"log"
	"reflect"
	"testing"
	"time"
)

type ABigStruct struct {
	Field0  string `db:"Field0"`
	Field1  string `db:"Field1"`
	Field2  string `db:"Field2"`
	Field3  string `db:"Field3"`
	Field4  string `db:"Field4"`
	Field5  string `db:"Field5"`
	Field6  string `db:"Field6"`
	Field7  string `db:"Field7"`
	Field8  string `db:"Field8"`
	Field9  string `db:"Field9"`
	Field10 string `db:"Field10"`
	Field11 string `db:"Field11"`
	Field12 string `db:"Field12"`
	Field13 string `db:"Field13"`
	Field14 string `db:"Field14"`
	Field15 string `db:"Field15"`
	Field16 string `db:"Field16"`
	Field17 string `db:"Field17"`
	Field18 string `db:"Field18"`
	Field19 string `db:"Field19"`
	Field20 string `db:"Field20"`
	Field21 string `db:"Field21"`
	Field22 string `db:"Field22"`
	Field23 string `db:"Field23"`
	Field24 string `db:"Field24"`
	Field25 string `db:"Field25"`
	Field26 string `db:"Field26"`
	Field27 string `db:"Field27"`
	Field28 string `db:"Field28"`
	Field29 string `db:"Field29"`
	Field30 string `db:"Field30"`
	Field31 string `db:"Field31"`
	Field32 string `db:"Field32"`
	Field33 string `db:"Field33"`
	Field34 string `db:"Field34"`
	Field35 string `db:"Field35"`
	Field36 string `db:"Field36"`
	Field37 string `db:"Field37"`
	Field38 string `db:"Field38"`
	Field39 string `db:"Field39"`
	Field40 string `db:"Field40"`
	Field41 string `db:"Field41"`
	Field42 string `db:"Field42"`
	Field43 string `db:"Field43"`
	Field44 string `db:"Field44"`
	Field45 string `db:"Field45"`
	Field46 string `db:"Field46"`
	Field47 string `db:"Field47"`
	Field48 string `db:"Field48"`
	Field49 string `db:"Field49"`
	Field50 string `db:"Field50"`
	Field51 string `db:"Field51"`
	Field52 string `db:"Field52"`
	Field53 string `db:"Field53"`
	Field54 string `db:"Field54"`
	Field55 string `db:"Field55"`
	Field56 string `db:"Field56"`
	Field57 string `db:"Field57"`
	Field58 string `db:"Field58"`
	Field59 string `db:"Field59"`
	Field60 string `db:"Field60"`
	Field61 string `db:"Field61"`
	Field62 string `db:"Field62"`
	Field63 string `db:"Field63"`
	Field64 string `db:"Field64"`
	Field65 string `db:"Field65"`
	Field66 string `db:"Field66"`
	Field67 string `db:"Field67"`
	Field68 string `db:"Field68"`
	Field69 string `db:"Field69"`
	Field70 string `db:"Field70"`
	Field71 string `db:"Field71"`
	Field72 string `db:"Field72"`
	Field73 string `db:"Field73"`
	Field74 string `db:"Field74"`
	Field75 string `db:"Field75"`
	Field76 string `db:"Field76"`
	Field77 string `db:"Field77"`
	Field78 string `db:"Field78"`
	Field79 string `db:"Field79"`
	Field80 string `db:"Field80"`
	Field81 string `db:"Field81"`
	Field82 string `db:"Field82"`
	Field83 string `db:"Field83"`
	Field84 string `db:"Field84"`
	Field85 string `db:"Field85"`
	Field86 string `db:"Field86"`
	Field87 string `db:"Field87"`
	Field88 string `db:"Field88"`
	Field89 string `db:"Field89"`
	Field90 string `db:"Field90"`
	Field91 string `db:"Field91"`
	Field92 string `db:"Field92"`
	Field93 string `db:"Field93"`
	Field94 string `db:"Field94"`
	Field95 string `db:"Field95"`
	Field96 string `db:"Field96"`
	Field97 string `db:"Field97"`
	Field98 string `db:"Field98"`
	Field99 string `db:"Field99"`
}

func TestCachedStructFields(t *testing.T) {
	a := ABigStruct{}
	rt := reflect.TypeOf(a)
	start := time.Now()
	fields, err := CachedStructFields(rt)
	costWithoutCache := time.Since(start)
	if err != nil {
		t.Error(err)
	}

	start = time.Now()
	fields, err = CachedStructFields(rt)
	costWithCache := time.Since(start)
	if err != nil {
		t.Error(err)
	}
	log.Printf("CachedStructFields costWithoutCache:%v costWithCache:%v", costWithoutCache, costWithCache)
	_ = fields
}

func TestCachedStructTags(t *testing.T) {
	a := ABigStruct{}
	rt := reflect.TypeOf(a)
	parse := func(tag string) (Tag, error) {
		return tag, nil
	}
	start := time.Now()
	tags, err := CachedStructTags(rt, "db", parse)
	costWithoutCache := time.Since(start)
	if err != nil {
		t.Error(err)
	}

	start = time.Now()
	tags, err = CachedStructTags(rt, "db", parse)
	costWithCache := time.Since(start)
	if err != nil {
		t.Error(err)
	}
	log.Printf("CachedStructTags costWithoutCache:%v costWithCache:%v", costWithoutCache, costWithCache)
	_ = tags
}
