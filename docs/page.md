# GitHub Wiki Markdown Sandbox

This page exists to test which Markdown features work in **GitHub Wiki**.

---

## 1. Headings
# H1
## H2
### H3
#### H4

---

## 2. Emphasis
**Bold**  
*Italic*  
***Bold + Italic***  
~~Strikethrough~~

---

## 3. Lists

### Unordered
- Item A
- Item B
  - Nested B.1
  - Nested B.2

### Ordered
1. First
2. Second
   1. Second.A
   2. Second.B

---

## 4. Links

### Root page link (works)
- [Home](Home)

### Subfolder link (works)
- [Docs index](docs/index)

### Subpage in subfolder (works)
- [Docs setup](docs/setup)

### Wiki-style link (root only)
- [[Home]]

### Wiki-style link (subfolder — should FAIL)
- [[docs/setup]]

---

## 5. Images

### Relative image (works if file exists)
![Example image](images/example.png)

### Absolute image (works)
![GitHub logo](https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png)

---

## 6. Code

### Inline
Use `git status` to check state.

### Fenced
```bash
git clone https://github.com/OWNER/REPO.wiki.git
