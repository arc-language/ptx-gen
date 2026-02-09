package codegen

import (
	"fmt"
	"strings"

	"github.com/arc-language/ptx-gen/builder"
)

// Emitter holds state during PTX text generation.
type Emitter struct {
    buf    strings.Builder
    indent int
}

// Emit takes a complete builder.Module and returns the PTX source string.
func Emit(mod *builder.Module) string {
    e := &Emitter{}
    e.emitModule(mod)
    return e.buf.String()
}

// --- Write helpers ---

// write appends raw text.
func (e *Emitter) write(s string) {
    e.buf.WriteString(s)
}

// writef appends formatted text.
func (e *Emitter) writef(format string, args ...interface{}) {
    e.buf.WriteString(fmt.Sprintf(format, args...))
}

// line writes an indented line followed by a newline.
func (e *Emitter) line(s string) {
    e.writeIndent()
    e.buf.WriteString(s)
    e.buf.WriteByte('\n')
}

// linef writes a formatted indented line.
func (e *Emitter) linef(format string, args ...interface{}) {
    e.writeIndent()
    e.buf.WriteString(fmt.Sprintf(format, args...))
    e.buf.WriteByte('\n')
}

// blank emits an empty line.
func (e *Emitter) blank() {
    e.buf.WriteByte('\n')
}

// writeIndent writes the current indentation (tab-based).
func (e *Emitter) writeIndent() {
    for i := 0; i < e.indent; i++ {
        e.buf.WriteByte('\t')
    }
}

// push increases indentation.
func (e *Emitter) push() {
    e.indent++
}

// pop decreases indentation.
func (e *Emitter) pop() {
    if e.indent > 0 {
        e.indent--
    }
}