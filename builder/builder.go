/* Here we have the builder design pattern,
which can help us construct a struct in a functional manner also given us
the chance of creating the said struct, "part-by-part"*/

package builder

type Employee struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

type Builder struct {
	e Employee
}

func (b *Builder) Build() Employee {
	return b.e
}

func (b *Builder) Name(name string) *Builder {
	b.e.Name = name
	return b
}

func (b *Builder) Role(role string) *Builder {
	if role == "CEO" {
		b.e.MaxSalary = 1_000_000
		b.e.MinSalary = 500_000
	} else {
		b.e.MaxSalary = 1_000
		b.e.MinSalary = 500
	}
	return b
}
