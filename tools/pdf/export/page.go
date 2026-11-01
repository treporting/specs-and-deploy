package export

import (
	"bytes"
	"html"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	gmhtml "github.com/yuin/goldmark/renderer/html"
)

// mermaid is the script of Mermaid, at a fixed version, loaded where a document holds a diagram; the
// hash makes Chrome refuse a script that differs from the one checked.
const mermaid = `<script src="https://cdn.jsdelivr.net/npm/mermaid@11.4.1/dist/mermaid.min.js"
  integrity="sha384-rbtjAdnIQE/aQJGEgXrVUlMibdfTSa4PQju4HDhN3sR2PmaKFzhEafuePsl9H/9I"
  crossorigin="anonymous" onerror="window.pdfError = 'the script of Mermaid did not load'"></script>
<script>
document.querySelectorAll('pre > code.language-mermaid').forEach(function (code) {
  var diagram = document.createElement('div');
  diagram.className = 'mermaid';
  diagram.textContent = code.textContent;
  code.parentElement.replaceWith(diagram);
});
if (window.mermaid) {
  mermaid.initialize({ startOnLoad: false });
  mermaid.run().then(function () { window.pdfReady = true; },
    function (e) { window.pdfError = 'Mermaid: ' + e; });
}
</script>`

// ready tells the printer that the page is done, where it holds no diagram.
const ready = `<script>window.pdfReady = true;</script>`

// style is the look of a page, after the one of GitHub: tables with borders, code on grey.
const style = `<style>
body { font-family: -apple-system, 'Segoe UI', Helvetica, Arial, sans-serif; font-size: 11px; line-height: 1.5; color: #1f2328; }
h1, h2 { border-bottom: 1px solid #d1d9e0; padding-bottom: .3em; }
table { border-collapse: collapse; margin: 1em 0; }
th, td { border: 1px solid #d1d9e0; padding: 4px 8px; text-align: left; vertical-align: top; }
th { background: #f6f8fa; }
tr { page-break-inside: avoid; }
code { font-family: ui-monospace, Menlo, Consolas, monospace; font-size: 90%; background: #f6f8fa; padding: .1em .3em; border-radius: 4px; }
pre { background: #f6f8fa; padding: 8px 12px; border-radius: 6px; overflow-x: hidden; white-space: pre-wrap; }
pre code { background: none; padding: 0; }
.mermaid { text-align: center; page-break-inside: avoid; }
</style>`

// markdown renders Markdown as GitHub does: its tables, strikethrough, task lists and links.
var markdown = goldmark.New(goldmark.WithExtensions(extension.GFM), goldmark.WithRendererOptions(gmhtml.WithUnsafe()))

// Page answers the HTML page of a document: its Markdown rendered, and its Mermaid diagrams drawn.
//
// @implements SR_69@1.1.1
func Page(title string, source []byte) (string, error) {
	var body bytes.Buffer
	if err := markdown.Convert(source, &body); err != nil {
		return "", err
	}
	script := ready
	if strings.Contains(body.String(), `class="language-mermaid"`) {
		script = mermaid
	}
	return Wrap(title, body.String()+script), nil
}

// Wrap answers a whole HTML page with its title, the style, and the body given.
func Wrap(title, body string) string {
	return "<!doctype html>\n<html><head><meta charset=\"utf-8\"><title>" + html.EscapeString(title) +
		"</title>\n" + style + "\n</head><body>\n" + body + "\n</body></html>\n"
}
