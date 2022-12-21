package weightconv

func KToL(k Kilogram) Pound { return Pound(k * 2.205) }
func LToK(l Pound) Kilogram { return Kilogram(l / 2.205) }
