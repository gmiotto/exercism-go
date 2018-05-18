package erratum

// Use 1
func Use(o ResourceOpener, input string) (err error) {
	res, err := o()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			if f, ok := r.(FrobError); ok {
				res.Defrob(f.defrobTag)
			}
			err = r.(error)
		}
		res.Close()
	}()

	res.Frob(input)
	return
}
