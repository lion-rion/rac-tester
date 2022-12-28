package tool

type Commad struct {
	Id          uint32
	Kickspeedx  float32
	Kickspeedz  float32
	Veltangent  float32
	Velnormal   float32
	Velangular  float32
	Spinner     bool
	Wheelsspeed bool
	Loop        int
	IsSim       bool
}
