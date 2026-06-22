# Contract: Wireframe render rules

- **id**: `spec:bpdsl.views.wireframe`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.views.overview`
- **contract_class**: `format`

## What this is

Render rules for the wireframe DSL — the `wireframe` section attached to a `state` node — output as HTML plus minimal CSS.

- **Target node**: only `type: state` nodes.
- **Input**: the single root wireframe tree under `state.wireframe`.
- **Output format**: HTML fragment + fixed CSS profile (the `.wf-*` namespace).
- **In scope**: screen structure, semantic containers, structural layout.
- **Out of scope**: visual style (color, font, border-radius, shadow, etc.), arbitrary CSS, direct class assignment, JS generation, data binding.

`layout` is limited to structural layout information needed to establish screen structure — visual style (color, font, border, radius, shadow) is not handled by the wireframe DSL.

> Source: V01-ADR-018, V01-ADR-019, V01-ADR-029, V01-ADR-042

## Current contract

### Element classification

Wireframe elements are classified as **container** or **leaf**.

#### container nodes

An element with children (`children`). `children` is required.

| type | description | HTML element |
|------|------|--------------|
| `col` | Vertical stack. | `div.wf-col` |
| `row` | Horizontal stack. | `div.wf-row` |
| `grid` | Grid layout. Column count via `cols` (required). | `div.wf-grid` |
| `card` | A semantic grouping. | `section.wf-card` |
| `sidebar` | Sidebar region. | `aside.wf-sidebar` |
| `header` | Header region. | `header.wf-header` |
| `footer` | Footer region. | `footer.wf-footer` |
| `main` | Main content region. | `main.wf-main` |

`main` is a container representing the primary content region, corresponding to `<main class="wf-main">` in HTML render.

#### leaf nodes

An element that cannot have `children`.

**interactive (may have `fires`)**

| type | description |
|------|------|
| `button` | Button. `label` required. |
| `input` | Text input. `label` required, `placeholder` optional. |
| `password` | Password input. `label` required, `placeholder` optional. |
| `select` | Dropdown. `label` required. |
| `checkbox` | Checkbox. `label` required. |
| `radio` | Radio button. `label` required. |

**non-interactive (no `fires`)**

| type | description |
|------|------|
| `text` | Text/label. `label` required. |
| `badge` | Badge/tag. `label` required. |
| `image` | Image placeholder. |
| `icon` | Icon placeholder. |
| `divider` | Divider line. |

### Field definitions

**Common to all elements**

| field | type | description |
|-----------|-----|------|
| `type` | string | Element kind (required). |
| `id` | string | Component ID. Output to HTML as `data-wf-id`, not the `id` attribute. |
| `label` | string | Display text. Meaning depends on `type`. Not needed for image/icon/divider/container nodes. |
| `layout` | object | Structural layout designation. Optional for all elements; which fields are allowed differs between container and leaf. |

**container-only**

| field | applies to | type | description |
|-----------|------|----|------|
| `children` | all containers | array | List of child nodes (required). |
| `cols` | grid | integer | Number of grid columns (required). |

**interactive-only**

| field | applies to | type | description |
|-----------|------|----|------|
| `fires` | button/input/password/select/checkbox/radio | string | Event ID fired on interaction. Output to HTML as `data-wf-fires`. |
| `disabled` | button/input/password/select/checkbox/radio | boolean | Inactive flag. |
| `placeholder` | input/password | string | Placeholder text for the input field. |

**for grid layout**

| field | type | description |
|-----------|-----|------|
| `span` | integer | How many grid columns to occupy. Written on a grid's child node. |

`span` only affects rendering as a direct child of `grid`. `span` outside `grid` has no meaning, but is not a parser error in v1.

### layout object

`layout` may be optionally specified on any wireframe element. In v1, the fields it handles are limited to structural layout.

```yaml
layout:
  width: 220
  height: 56
  min_width: 120
  min_height: 80
  grow: true
  gap: 16
  padding: 16
  align: center
  justify: between
  scroll: y
```

#### layout field catalog

| field | type | applies to | content |
|---|---|---|---|
| `width` | size | all elements | Width. Numbers are px. Reserved words `fill` / `fit` allowed. |
| `height` | size | all elements | Height. Numbers are px. Reserved words `fill` / `fit` allowed. |
| `min_width` | int | all elements | Minimum width in px. `fill` / `fit` not allowed. |
| `min_height` | int | all elements | Minimum height in px. `fill` / `fit` not allowed. |
| `grow` | bool | direct child of row/col | When `true`, occupies the remaining space. |
| `gap` | int | container | Gap in px between children. |
| `padding` | int or object | container | Structural inner padding of the container. Not supported on leaf in v1. |
| `align` | enum | container | Cross-axis alignment: `start` / `center` / `end` / `stretch`. |
| `justify` | enum | container | Main-axis alignment: `start` / `center` / `end` / `between`. |
| `scroll` | enum | container | `none` / `x` / `y` / `both`. |

#### size values

`width` / `height` accept one of:

| value | meaning | HTML/CSS conversion |
|---|---|---|
| number | Fixed px value. | `Npx` |
| `fill` | 100% of the parent's available size. | `width: 100%` / `height: 100%` |
| `fit` | Content size. | `fit-content` for width; `auto` for height. |

```yaml
layout:
  width: 220       # width: 220px
  height: 56       # height: 56px; min-height: 56px
```

A fixed `height` converts to both `height` and `min-height`, to preserve vertical occupied space for header/footer/image etc. A fixed `width` does not auto-add `min-width` — specify `min_width` explicitly if needed.

```yaml
layout:
  width: fill      # width: 100%
  height: fit      # height: auto
```

Numbers are always treated as px. Arbitrary CSS strings like `"80%"` / `"12rem"` / `"calc(...)"` are not allowed.

`min_width` / `min_height` accept px integers only — `fill` / `fit` are not allowed.

#### padding object

`padding` may only be specified on a container. It accepts a number or an object.

```yaml
layout:
  padding: 16
```

```yaml
layout:
  padding:
    x: 16
    y: 12
```

```yaml
layout:
  padding:
    top: 8
    right: 16
    bottom: 8
    left: 16
```

`x` is shorthand for left/right; `y` is shorthand for top/bottom. Resolution order: initialize `top` / `right` / `bottom` / `left` to 0, apply `x` to left/right if present, apply `y` to top/bottom if present, then apply any individual `top` / `right` / `bottom` / `left` overrides last.

```yaml
layout:
  padding:
    x: 16
    top: 8
```

The above converts to `padding: 8px 16px 0px 16px`.

Leaf-element padding is not handled in v1 — it tends to become a button/input/text visual-tuning concern, blurring the boundary between structural layout and visual style.

#### fill vs. grow

`width: fill` / `height: fill` means `100%` relative to the parent's available size. It does not mean occupying flex remaining space.

`grow: true` means occupying the remaining space as a child of `row` / `col`. The base HTML/CSS conversion is `flex: 1 1 0%; min-width: 0; min-height: 0`.

```yaml
- type: row
  children:
    - type: sidebar
      layout:
        width: 220
      children: [...]
    - type: main
      layout:
        grow: true
        scroll: y
      children: [...]
```

To make an element take up the screen's remaining space, use `grow: true`, not `fill`. Specifying `width: fill` / `height: fill` on a direct child of `row` / `col` does not produce flex remaining-space occupation. In particular, specifying `width: fill` on a direct child of `row` risks pushing out sibling elements — use `grow: true` for remaining-space occupation instead.

### Root structure

Directly under `wireframe:` is a **single container node**. Multiple roots are not allowed.

```yaml
# OK
wireframe:
  type: col
  children: [...]

# NG (multiple roots)
wireframe:
  - type: col
    children: [...]
  - type: col
    children: [...]
```

### Per-state display differences

States like `loading` / `error` are each defined as a **separate state node**, each with its own `wireframe`. There is no display-condition field (e.g. `visible_when`) within a single `wireframe`.

```yaml
# OK: separate state nodes
- id: login_screen
  type: state
  wireframe:
    type: col
    children:
      - type: button
        label: Log in
        fires: submit_clicked

- id: login_loading
  type: state
  wireframe:
    type: col
    children:
      - type: text
        label: Submitting...
      - type: button
        label: Log in
        disabled: true

# NG: conditionally branching state within wireframe
wireframe:
  type: col
  children:
    - type: text
      label: Submitting...
      state: login_loading   # this field does not exist
```

### HTML/CSS render profile

The HTML renderer outputs the wireframe tree as an HTML fragment. It does not generate `DOCTYPE`, `html`, `head`, or `body`.

#### Output contract

- Every wireframe element gets a `.wf-*` namespaced class.
- YAML `id` is not output to HTML's `id` attribute — it goes to `data-wf-id`. This avoids HTML `id` collisions when multiple wireframe fragments are shown on the same page.
- YAML `fires` is output to `data-wf-fires`.
- `label` / `placeholder` / `id` / `fires` are HTML-escaped.
- No JS is generated.
- Arbitrary CSS is not accepted.
- `layout` converts to deterministic inline style. Even if an implementation uses fixed utility classes, they must convert to the same meaning.

#### HTML element mapping

| type | output |
|---|---|
| `col` | `<div class="wf-col">` |
| `row` | `<div class="wf-row">` |
| `grid` | `<div class="wf-grid">` |
| `card` | `<section class="wf-card">` |
| `sidebar` | `<aside class="wf-sidebar">` |
| `header` | `<header class="wf-header">` |
| `footer` | `<footer class="wf-footer">` |
| `main` | `<main class="wf-main">` |
| `button` | `<button class="wf-button">` |
| `input` | `<div class="wf-field"><label>...</label><input type="text" /></div>` |
| `password` | `<div class="wf-field"><label>...</label><input type="password" /></div>` |
| `select` | `<div class="wf-field"><label>...</label><select></select></div>` |
| `checkbox` | `<label class="wf-checkbox"><input type="checkbox" /> ...</label>` |
| `radio` | `<label class="wf-radio"><input type="radio" /> ...</label>` |
| `text` | `<span class="wf-text">...</span>` |
| `badge` | `<span class="wf-badge">...</span>` |
| `image` | `<div class="wf-image">[image]</div>` |
| `icon` | `<span class="wf-icon">[icon]</span>` |
| `divider` | `<hr class="wf-divider" />` |

#### fixed CSS profile

The renderer may assume the following fixed CSS as meaning. The reference CSS file is `wireframe.css` (in this same `views/` directory). The excerpt below shows the main rules; the full reference CSS in that file is authoritative.

```css
.wf-col { display: flex; flex-direction: column; gap: 8px; }
.wf-row { display: flex; flex-direction: row; gap: 8px; }
.wf-grid { display: grid; gap: 8px; }
.wf-header, .wf-footer, .wf-sidebar, .wf-main, .wf-card { display: flex; flex-direction: column; gap: 8px; }
.wf-field { display: flex; flex-direction: column; gap: 4px; }
```

This is the render profile's own fixed CSS, not a style arbitrarily specified from YAML. `wireframe.css` contains structural CSS only — no color, border, radius, shadow, font, background-color, or preview-only decoration.

#### preview CSS

The CSS file for visual confirmation is `wireframe.preview.css` (in this same `views/` directory).

`wireframe.preview.css` is auxiliary CSS to make the HTML fragment easier to check in a browser — it is not part of the wireframe DSL's semantics. It may include preview-only visual style: thin borders, background colors, rounded corners, minimal button/input/select appearance, etc.

The Go renderer's HTML fragment output and golden-test expected values do not depend on `wireframe.preview.css`.

#### layout conversion

| YAML | HTML/CSS meaning |
|---|---|
| `layout.width: 220` | `width: 220px` |
| `layout.width: fill` | `width: 100%` |
| `layout.width: fit` | `width: fit-content` |
| `layout.height: 56` | `height: 56px; min-height: 56px` |
| `layout.height: fill` | `height: 100%` |
| `layout.height: fit` | `height: auto` |
| `layout.min_width: 120` | `min-width: 120px` |
| `layout.min_height: 80` | `min-height: 80px` |
| `layout.grow: true` | `flex: 1 1 0%; min-width: 0; min-height: 0` |
| `layout.gap: 16` | `gap: 16px` |
| `layout.padding: 16` | `padding: 16px` |
| `layout.padding.x: 16` + `layout.padding.y: 8` | `padding: 8px 16px 8px 16px` |
| `layout.padding.top/right/bottom/left` | `padding: {top}px {right}px {bottom}px {left}px`. Resolution order: `0` init → apply `x` / `y` → override with individual sides. |
| `layout.align: start` | `align-items: flex-start` |
| `layout.align: center` | `align-items: center` |
| `layout.align: end` | `align-items: flex-end` |
| `layout.align: stretch` | `align-items: stretch` |
| `layout.justify: start` | `justify-content: flex-start` |
| `layout.justify: center` | `justify-content: center` |
| `layout.justify: end` | `justify-content: flex-end` |
| `layout.justify: between` | `justify-content: space-between` |
| `layout.scroll: none` | `overflow: visible` |
| `layout.scroll: x` | `overflow-x: auto` |
| `layout.scroll: y` | `overflow-y: auto` |
| `layout.scroll: both` | `overflow: auto` |

When `grow: true` and `min_width` / `min_height` are both specified, the explicit `min_width` / `min_height` takes precedence.

#### grid render

`grid`'s base display comes from the fixed CSS profile's `.wf-grid { display: grid; ... }`. `grid.cols` converts to `grid-template-columns`, output as inline style. `display: grid` is not duplicated into inline style.

```css
grid-template-columns: repeat({cols}, 1fr)
```

A child node with `span` gets the following added to its own style:

```css
grid-column: span {span}
```

### Render example

**A screen with `main` + `layout`**

YAML:

```yaml
- id: cart
  type: state
  wireframe:
    type: col
    children:
      - type: header
        layout:
          height: 56
        children:
          - type: text
            label: Shopping Cart
      - type: row
        layout:
          grow: true
        children:
          - type: sidebar
            layout:
              width: 220
            children:
              - type: text
                label: Order Summary
          - type: main
            layout:
              grow: true
              scroll: y
            children:
              - type: text
                label: Cart Items
```

HTML output:

```html
<div class="wf-col">
  <header class="wf-header" style="height: 56px; min-height: 56px;">
    <span class="wf-text">Shopping Cart</span>
  </header>
  <div class="wf-row" style="flex: 1 1 0%; min-width: 0; min-height: 0;">
    <aside class="wf-sidebar" style="width: 220px;">
      <span class="wf-text">Order Summary</span>
    </aside>
    <main class="wf-main" style="flex: 1 1 0%; min-width: 0; min-height: 0; overflow-y: auto;">
      <span class="wf-text">Cart Items</span>
    </main>
  </div>
</div>
```

## Rules

- `layout` field applicability (container-only vs. all-elements vs. row/col-direct-child-only) follows §layout field catalog above; the renderer must not silently accept a field outside its applicable scope.
- `wireframe.css` defines structural CSS only; visual style belongs exclusively to `wireframe.preview.css`, and golden-test expected output never depends on the preview file.

## Validation rules

The parser/validator checks the following:

- Directly under `wireframe:` is a single container node. Multiple roots are invalid.
- A container node requires `children`.
- A leaf node cannot have `children`.
- `main` is a container, so it requires `children`.
- `grid` requires `cols`; `cols` must be an integer ≥ 1.
- `layout` is optional on all elements.
- `width` / `height` / `min_width` / `min_height` may be specified on either container or leaf.
- `grow` may be written on either container or leaf, but `grow: true` only takes effect as a direct child of `row` / `col`.
- `gap` / `padding` / `align` / `justify` / `scroll` are valid only on a container. Specifying them on a leaf is a parser error.
- Specifying `grow: true` on anything other than a direct child of `row` / `col` is a parser error.
- An undefined field under `layout` is a parser error.
- Fields that directly specify HTML/CSS implementation detail (`style` / `class` / `css` etc.) are a parser error.
- `fires` may only be specified on an interactive leaf.
- `id` / `label` / `placeholder` / `fires` are escaped on HTML output.

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.views.overview` | Parent overview; view kind catalog. |
| `spec:bpdsl.dsl.nodes.application` | `state` node definition that `wireframe` attaches to. |
