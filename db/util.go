package db

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mmcloughlin/meow"
)

func (r AccessRule) PolicyLine() string {
	vList := []string{}
	vs := []pgtype.Text{r.V0, r.V1, r.V2, r.V3, r.V4, r.V5}

	for _, v := range vs {
		if v.Valid {
			vList = append(vList, v.String)
		}
	}

	if len(vList) == 0 {
		return r.Ptype.String
	}

	return fmt.Sprintf(
		"%s, %s",
		r.Ptype.String,
		strings.Join(vList, ", "),
	)
}

func (r *AccessRule) SetV(value string, i int) {
	switch i {
	case 0:
		r.V0 = text(value)
	case 1:
		r.V1 = text(value)
	case 2:
		r.V2 = text(value)
	case 3:
		r.V3 = text(value)
	case 4:
		r.V4 = text(value)
	case 5:
		r.V5 = text(value)
	default:
		panic("invalid v index")
	}
}

func (r *AccessRule) Scan(ptype string, rule []string) {
	r.Ptype = text(ptype)

	for i, v := range rule {
		r.SetV(v, i)
	}

	r.ID = r.GetID()
}

func (r AccessRule) GetID() string {
	rule := []string{r.V0.String, r.V1.String, r.V2.String, r.V3.String, r.V4.String, r.V5.String}

	data := strings.Join(append([]string{r.Ptype.String}, rule...), ",")
	sum := meow.Checksum(0, []byte(data))
	return fmt.Sprintf("%x", sum)
}

func (r AccessRule) String() string {
	rule := []string{r.V0.String, r.V1.String, r.V2.String, r.V3.String, r.V4.String, r.V5.String}
	rule = slices.DeleteFunc(rule, func(v string) bool {
		return len(v) == 0
	})

	return fmt.Sprintf("%s, %s", r.Ptype.String, strings.Join(rule, ", "))
}

func text(s string) pgtype.Text {
	if len(s) == 0 {
		return pgtype.Text{
			Valid: false,
		}
	}
	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}
