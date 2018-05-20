package main

import (
    "fmt"
    "os"
    "math"
    "image"
    "image/color"
    "image/png"
    "flag"
)

func debug(args... interface{}) {
    fmt.Fprintln(os.Stderr,args...)
}

const inv64max = 1./(1<<64)

// xorshift64 //
func rng64(seed uint64) float64 {
    x := seed ^ 0x91039276
    x ^= x >> 12; // a
    x ^= x << 25; // b
    x ^= x >> 27; // c
    return float64(x * 0x2545F4914F6CDD1D)*inv64max;
}

func rngf1(x float64) float64 {
    seed := math.Float64bits(x*float64(3.42536932593826360421057328563))
    return 2.*rng64(seed)-1
}

func rngf2(x,y float64) float64 {
    t := math.Float64bits(x*float64(2.4850375104769257201575375368361538)) ^
         math.Float64bits(y*float64(2.3867397636247669375625836582310473))
    seed := t
    return 2.*rng64(seed)-1
}

func perlin_noise_1d(x, seed float64) float64 {
    xi, xf := math.Modf(x)
    v00x := rngf2(xi,seed)
    v10x := rngf2(xi+1,seed)
    ux := xf*xf*(3.0-2.0*xf)
    c00 := v00x*xf
    c10 := v10x*(xf-1)
    d1 := c00*(1-ux)+c10*ux
    r := d1
    return r
}

// func perlin_noise_2d(x,y float64) float64 {
//     xi, xf := math.Modf(x)
//     yi, yf := math.Modf(y)
//     v00x := rngf2(xi,yi)
//     v00y := rngf2(xi+0.1,yi)
//     v10x := rngf2(xi+1,yi)
//     v10y := rngf2(xi+1+0.1,yi)
//     v01x := rngf2(xi,yi+1)
//     v01y := rngf2(xi+0.1,yi+1)
//     v11x := rngf2(xi+1,yi+1)
//     v11y := rngf2(xi+1+0.1,yi+1)
//     ux := xf*xf*(3.0-2.0*xf)
//     uy := yf*yf*(3.0-2.0*yf)
//     c00 := v00x*xf+v00y*yf
//     c10 := v10x*(xf-1)+v10y*yf
//     c01 := v01x*xf+v01y*(yf-1)
//     c11 := v11x*(xf-1)+v11y*(yf-1)
//     d1 := c00*(1-ux)+c10*ux
//     d2 := c01*(1-ux)+c11*ux
//     r := d1*(1-uy)+d2*uy
//     return r
// }


func dist_width(dx,dy float64) float64 {
    var oy float64
    if dy >=0 {
        if dy <= 1 {
            oy = dy
        } else {
            oy = 1
        }
    } else {
        oy = 0
    }
    w := shape_width(oy,rml,shapel)
    return (dx*dx+(dy-oy)*(dy-oy)) / (w*w)
}

func branch_scale(offset float64) float64 {
    return 1/(shape_width(offset,rmb,shapeb)*sb)
}

func shape_width(offset,end,shape float64) float64 {
    if shape < 0 {
        if end > 0.99999 {
            return 1
        }
        k := 1-end
        delta := k-shape
        b := 0.5*(delta*delta/k-k)
        do := offset*delta
        return 1+b-math.Sqrt(b*b+do*do)
    } else if shape < 1e-5 {
        return (1-offset) + offset*end
    } else if shape > 0 {
        a := (end-1)/(math.Exp(-shape)-1)
        return a*(math.Exp(-offset*shape)-1)+1
    }
    return 1
}

func noise_offset(pos, seed float64) float64 {
    res := 0.
        freq := 1.2
        k := 1.
        for i:=0; i<noise_dim; i++ {
        res += perlin_noise_1d(pos*freq,seed+0.1+float64(i))*k
        freq *= noise_lacunarity
        k *= noise_gain
    }
    return res*noise_amount
    //return perlin_noise_1d(pos*1,seed)*0.2+perlin_noise_1d(pos*3,seed+50)*0.15+perlin_noise_1d(pos*10,seed+100)*0.1
}

func iter_branch(dx,dy,ybo float64, right bool, bl int, seed float64) float64 {
    d := math.Inf(0)
    dic := di
    for dyo:=0.;;dyo+=dic {
        yb := dyo+ybo*dic + rngf2(dyo,seed) * branch_start_jitter * dic
        if yb > dm {
            break
        }
        tx, ty := dcoss*(1-yb)+dcose*yb, dsins*(1-yb)+dsine*yb
        if right {
            tx = -tx
        }
        s := branch_scale(yb)
        dxc, dyc := dx,dy-yb
        dxcr1, dycr1 := (dxc*ty-dyc*tx)*s,(dyc*ty+dxc*tx)*s
        d1 := iter(dxcr1,dycr1,bl-1,seed+1+dyo)*is
        if d1 < d {
            d = d1
        }
        dic *= sl
    }
    return d
}

func iter(dx,dy float64, bl int, seed float64) float64 {
    if dx > dcutx || dx < -dcutx ||
        dy < -dcuty || dy-1 > dcuty {
        return math.Inf(0)
    }
    evals++
    dyn := dy
    if dyn < 0 {
        dyn = 0
    }
    if dyn > 1 {
        dyn = 1
    }
    dx += noise_offset(dyn,seed)
    d := dist_width(dx,dy)
    if bl != 0 {
        dc := iter_branch(dx,dy,dor,false,bl,seed)
        if dc < d {
            d = dc
        }
        dc = iter_branch(dx,dy,dol,true,bl,seed+0.123426262235623723)
        if dc < d {
            d = dc
        }
    }
    return d
}

func create_image() {
    var error error
    var file *os.File
    if filename != "" && filename != "" {
        file, error = os.Create(filename)
    } else {
        file = os.Stdout
    }
    if error != nil {
        fmt.Fprintln(os.Stderr,"error opening file:",error)
        return
    }
    result := image.NewRGBA( image.Rectangle{image.Point{0,0},image.Point{width,height}} )
    for y := 0; y<height; y++ {
        for x := 0; x<width; x++ {
            v := iter(float64(x-width/2)*scalex,float64(y-height/8)*scaley,bd,base_seed)
            v = 1/(1+v*100)
            c := uint8(v*0xff)
            result.Set(x,width-1-y,color.RGBA{c,c,c,0xff})
        }
    }
    png.Encode( file, result )
    fmt.Fprintln(os.Stderr,"branch evaluations:",evals)
}

var sb float64
var is float64
var rs float64
var re float64
var di float64
var dor float64
var dol float64
var dm float64
var sl float64
var dcutx float64
var dcuty float64
var rml float64
var rmb float64
var shapel float64
var shapeb float64
var bd int
var noise_dim int
var noise_lacunarity float64
var noise_gain float64
var noise_amount float64
var branch_start_jitter float64
var base_seed float64

var dcoss, dsins float64
var dcose, dsine float64

var width int
var height int

var scalex float64
var scaley float64

var filename string

var evals int

func validate_arg_f64(val float64, name string, min,max float64) bool {
    if val < min || val > max {
        fmt.Fprintf(os.Stderr,"error: %s must be between %f and %f\n",name,min,max)
        return false
    }
    return true
}

func validate_arg_lopen_f64(val float64, name string, min,max float64) bool {
    if val <= min || val > max {
        fmt.Fprintf(os.Stderr,"error: %s must be between %f and %f\n",name,min,max)
        return false
    }
    return true
}

func validate_arg_lopen_pos_f64(val float64, name string) bool {
    if val <= 0 {
        fmt.Fprintf(os.Stderr,"error: %s must be > 0\n",name)
        return false
    }
    return true
}

func validate_arg_pos_f64(val float64, name string) bool {
    if val < 0 {
        fmt.Fprintf(os.Stderr,"error: %s must be >= 0\n",name)
        return false
    }
    return true
}

func validate_arg_pos_int(val int, name string) bool {
    if val < 0 {
        fmt.Fprintf(os.Stderr,"error: %s must be >= 0\n",name)
        return false
    }
    return true
}

func validate_arg_gt_int(val int, name string, gt int) bool {
    if val < gt {
        fmt.Fprintf(os.Stderr,"error: %s must be >= %i\n",name,gt)
        return false
    }
    return true
}

func validate_sl(sl, di, dm float64) bool {
    if ! validate_arg_lopen_f64(sl,"sl",0,1) {
        return false
    }
    sl_min := dm/(dm+di)
    if sl <= sl_min {
        fmt.Fprintf(os.Stderr,"error: sl must be >= dm*(dm+di) = %f\n",sl_min)
        //return false
    }
    return true
}

func main() {
    flag.Float64Var(&sb,"sb",0.6,"scale (relative size) of branch")
    flag.Float64Var(&is,"is",1.1,"intensity scale factor of branch (higher values make more deeply nested branches finer)")
    flag.Float64Var(&rs,"rs",0.3,"rotation of branch (near root of parent) in multiples of 180 degrees")
    flag.Float64Var(&re,"re",0.2,"rotation of branch (near tip of parent) in multiples of 180 degrees")
    flag.Float64Var(&di,"di",0.55,"initial relative distance of branches on same parent (near root of parent). Use -sl to make branches denser near tip.")
    flag.Float64Var(&sl,"sl",0.7,"distance scale of branches at same parent (lower values make branches denser near tip of parent). Note: if too low, number of branches goes to infinity, resuning in an infinite loop (sl must be > dm*/(1-dl))")
    flag.Float64Var(&dol,"dol",0.4,"distance of first left branch from root of parent (in multiples of parent length)")
    flag.Float64Var(&dor,"dor",0.6,"distance of first right branch from root of parent (in multiples of parent length)")
    flag.Float64Var(&dm,"dm",0.9,"maximum distance of branches from root of parent (in multiples of parent length)")
    flag.Float64Var(&rml,"rml",0.1,"relative width of branch at tip")
    flag.Float64Var(&rmb,"rmb",0.1,"relative scale of branches starting near tip of parent. Note: if significantly different from -rml, result may show visible discontinuities.")
    flag.Float64Var(&shapel,"shapel",0,"shape of branch (positive values make branches more pointy, negative ones more blunted)")
    flag.Float64Var(&shapeb,"shapeb",0,"same as -shapel but for -rmb. Note: if significantly different from -rml, result may show visible discontinuities.")
    flag.IntVar(&bd,"bd",3,"recursion depth (number of nested branches).")
    flag.Float64Var(&branch_start_jitter,"bj",0.1,"Branch jitter (random displacement of branches along parent)")
    flag.IntVar(&noise_dim,"nd",3,"dimension of fractal noise (higher values will produce larger spectrum)")
    flag.Float64Var(&noise_lacunarity,"nl",1.7,"lacunarity of fractal noise (higher values will stretch spectrum to higher frequencies)")
    flag.Float64Var(&noise_gain,"ng",0.5,"gain of noise (higher values will make high-frequency components more prominent)")
    flag.Float64Var(&noise_amount,"na",0.1,"amount of noise (higher values will result make more noisy, zero will produce completely straight branches)")
    flag.Float64Var(&dcutx,"cx",2,"iteration cutoff distance in direction perpendicular to branch (this is a performance optimization parameter, higher values make computation take longer, but if too low, branches will be truncated, leading to ugly artifacts).")
    flag.Float64Var(&dcuty,"cy",0.5,"iteration cutoff distance in direction along branch (this is a performance optimization parameter, higher values make computation take longer, but if too low, branches will be truncated, leading to ugly artifacts).")
    flag.IntVar(&width,"width",512,"width of output image in pixels")
    flag.IntVar(&height,"height",512,"height of output image in pixels")
    flag.StringVar(&filename,"f","tree.png","output filename (use empty string or '-' for stdout)")
    flag.Float64Var(&base_seed,"seed",0,"random seed (changing this will produce different random values)")

    flag.Parse()

    if !validate_arg_f64(sb,"sb",0,1) ||
        !validate_arg_lopen_pos_f64(is,"is") ||
        !validate_arg_f64(rs,"rs",0,1) ||
        !validate_arg_f64(re,"re",0,1) ||
        !validate_arg_lopen_f64(di,"di",0,1) ||
        !validate_arg_f64(dm,"dm",0,1) ||
        !validate_sl(sl,di,dm) ||
        !validate_arg_f64(dol,"dol",0,1) ||
        !validate_arg_f64(dor,"dor",0,1) ||
        !validate_arg_lopen_f64(rml,"rml",0,1) ||
        !validate_arg_lopen_f64(rmb,"rmb",0,1) ||
        !validate_arg_pos_int(bd,"bd") ||
        !validate_arg_f64(branch_start_jitter,"bj",0,1) ||
        !validate_arg_pos_int(noise_dim,"nd") ||
        !validate_arg_lopen_pos_f64(noise_lacunarity,"nl") ||
        !validate_arg_pos_f64(noise_gain,"ng") ||
        !validate_arg_pos_f64(noise_amount,"na") ||
        !validate_arg_lopen_pos_f64(dcutx,"cx") ||
        !validate_arg_lopen_pos_f64(dcuty,"cy") ||
        !validate_arg_gt_int(width,"width",16) ||
        !validate_arg_gt_int(height,"height",16) {
        os.Exit(1)
    }

    scalex = 1.3/float64(width)
    scaley = 1.3/float64(height)
    rs *= math.Pi
    re *= math.Pi

    dcoss, dsins = math.Sincos( rs )
    dcose, dsine = math.Sincos( re )

    create_image()
}
