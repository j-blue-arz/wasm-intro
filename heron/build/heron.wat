(module
  (type (;0;) (func (result i32)))
  (type (;1;) (func))
  (type (;2;) (func (param i32)))
  (type (;3;) (func (param i32) (result i32)))
  (type (;4;) (func (param i32) (result f64)))
  (type (;5;) (func (param f64) (result f64)))
  (func (;0;) (type 1)
    nop)
  (func (;1;) (type 5) (param f64) (result f64)
    (local f64 i32)
    f64.const 0x1.5p+5 (;=42;)
    local.set 1
    i32.const 1
    local.set 2
    loop  ;; label = @1
      local.get 1
      local.get 0
      local.get 1
      f64.div
      f64.add
      f64.const 0x1p-1 (;=0.5;)
      f64.mul
      local.set 1
      local.get 2
      i32.const 1000
      i32.eq
      i32.eqz
      if  ;; label = @2
        local.get 2
        i32.const 1
        i32.add
        local.set 2
        br 1 (;@1;)
      end
    end
    local.get 1)
  (func (;2;) (type 4) (param i32) (result f64)
    local.get 0
    f64.convert_i32_s
    call 1)
  (func (;3;) (type 0) (result i32)
    global.get 0)
  (func (;4;) (type 2) (param i32)
    local.get 0
    global.set 0)
  (func (;5;) (type 3) (param i32) (result i32)
    global.get 0
    local.get 0
    i32.sub
    i32.const -16
    i32.and
    local.tee 0
    global.set 0
    local.get 0)
  (func (;6;) (type 0) (result i32)
    i32.const 1024)
  (table (;0;) 2 2 funcref)
  (memory (;0;) 256 256)
  (global (;0;) (mut i32) (i32.const 5243920))
  (export "memory" (memory 0))
  (export "heron" (func 1))
  (export "heron_int" (func 2))
  (export "_initialize" (func 0))
  (export "__indirect_function_table" (table 0))
  (export "__errno_location" (func 6))
  (export "stackSave" (func 3))
  (export "stackRestore" (func 4))
  (export "stackAlloc" (func 5))
  (elem (;0;) (i32.const 1) func 0))
