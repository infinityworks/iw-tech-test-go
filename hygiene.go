package hygiene

type Authority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AuthorityRating struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type FSAAuthorities struct {
	Authorities []FSAAuthority `json:"authorities"`
}

type FSAAuthority struct {
	ID   int    `json:"LocalAuthorityId"`
	Name string `json:"Name"`
}
