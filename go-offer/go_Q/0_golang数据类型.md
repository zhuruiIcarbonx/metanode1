Go 语言的数据类型可以分为**基本类型**、**复合类型**（或叫复杂类型）、**引用类型**和**接口类型**等。以下是详细的分类及说明：

---

### **1. 基本类型（Primitive Types）**
包括简单的数值、布尔值和字符串：
- **布尔型**：`bool`（`true`/`false`）
- **整数型**：
  - 有符号：`int`（系统决定位数）、`int8`、`int16`、`int32`（`rune`）、`int64`
  - 无符号：`uint`、`uint8`（`byte`）、`uint16`、`uint32`、`uint64`、`uintptr`
- **浮点型**：`float32`、`float64`
- **复数型**：`complex64`、`complex128`
- **字符串**：`string`（不可变 UTF-8 序列）

> 📌 `byte` 是 `uint8` 的别名，`rune` 是 `int32` 的别名（表示 Unicode 码点）。

---

### **2. 复合类型（Composite Types）**
通过组合基本类型构建的复杂类型：
- **数组（Array）**：固定长度，如 `[3]int{1, 2, 3}`
- **结构体（Struct）**：自定义字段集合，如：
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

---

### **3. 引用类型（Reference Types）**
底层指向数据结构的指针，包括：
- **切片（Slice）**：动态数组，基于数组封装（如 `[]int`）
- **映射（Map）**：键值对集合，如 `map[string]int`
- **通道（Channel）**：用于协程通信，如 `chan int`
- **函数（Function）**：函数类型（一等公民），如 `func(int) int`
- **指针（Pointer）**：存储内存地址，如 `*int`

> ⚠️ 引用类型的零值为 `nil`。

---

### **4. 接口类型（Interface Types）**
定义行为契约的抽象类型：
- **接口（Interface）**：一组方法签名，例如：
  ```go
  type Writer interface {
      Write([]byte) (int, error)
  }
  ```
- 特殊接口：
  - `error`：内置错误接口。
  - `any`（Go 1.18+）：空接口的别名（相当于 `interface{}`）。

---

### **5. 其他类型**
- **类型别名（Type Aliases）**：通过 `type` 创建别名，如 `type MyInt = int`
- **自定义类型（Defined Types）**：基于现有类型创建新类型，如 `type MyInt int`

---

### **总结表格**
| **类别**       | **类型举例**                          | **特点**                     |
|----------------|--------------------------------------|-----------------------------|
| 基本类型       | `int`, `string`, `bool`             | 值类型，直接存储数据         |
| 复合类型       | `[3]int`, `struct{}`                | 值类型，组合基本类型         |
| 引用类型       | `[]int`, `map`, `chan`, `func`      | 底层指针，零值为 `nil`       |
| 接口类型       | `error`, `interface{}`              | 抽象行为，隐式实现           |
| 自定义类型     | `type MyInt int`                    | 增强类型安全性               |

---

### **示例代码**
```go
var (
    n int = 42                      // 基本类型
    arr = [3]int{1, 2, 3}          // 数组（复合类型）
    s = []int{1, 2}                // 切片（引用类型）
    m = map[string]int{"a": 1}     // 映射（引用类型）
    ch = make(chan int)            // 通道（引用类型）
    w io.Writer                     // 接口类型
)
```

如果需要进一步探讨某类类型的具体用法或场景，可以随时提问！  
（当前时间：2025-10-09 11:01:31，星期四）
