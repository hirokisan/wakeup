package store

// RelationalOperator :
type RelationalOperator string

// TODO : 他にもいろいろ比較演算ある
const (
	// RelationalOperatorEqual :
	RelationalOperatorEqual = RelationalOperator("equal")
)

// Query :
// TODO : 論理演算もある
// TODO : sort, offsetとかもある
type Query struct {
	Field    string
	Operator RelationalOperator
	Value    string // TODO : 良い感じに型を伝達させると良さそう
}
