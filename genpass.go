package spg

/**
 * Secure Password Generator.
 *
 * Lovingly handcrafted by Dave Teare on October 20th, 2017
 * Cruelly butchered by Jeffrey Goldberg in February 2018
 */

/** Attributes

	Attributes are all of the knobs that can be set for specifying details of what
	of the password requirements.

	There are two kinds of attributes (other than the Length). There are those that are
	specific for character type passwords and those that are relevant to things based
	on word lists. Ultimately we are sticking this all in the same type so that we
	can have a common Generator interface. So this gets kind of messy
 **/

// Password is what gets generated by the Generators
type Password struct {
	Tokens  []Token
	Entropy float32 // Entropy in bits of the Recipe from which this password was generated
}

// Generator is a fully configured password recipe
type Generator interface {
	Generate() (*Password, error)
	Entropy() float32
}

// Atoms returns the tokens (words, syllables, characters) that compromise the bulk of the password
func (p Password) Atoms() []string { return p.tokensOfType(AtomTokenType) }

// Separators are the separators between tokens.
// If this list is shorter than one less then the number of tokens, the last separator listed
// is used repeatedly to separate subsequent tokens.
// If this is nil, it is taken as nil no separators between tokens
func (p Password) Separators() []string { return p.tokensOfType(SeparatorTokenType) }

// String is the Stringer. It produces the password as string one might expect
func (p Password) String() string {
	pw := ""
	for _, tok := range p.Tokens {
		pw += tok.Value
	}
	return pw
}
