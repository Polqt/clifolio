
✅ 2. Set up cmd/portfolio/main.go

Add CLI flags (debug, theme).

Load theme.

Initialize the root Bubble Tea app model.

Run Bubble Tea with fullscreen mode.

❗ Goal of Phase 1

You have a runnable “blank” TUI that opens to an empty screen.

PHASE 2 — Core UI System (Your App Skeleton)

✅ 3. Build the Root App Model

Create application state enum (Intro, Menu, Projects, Skills, Experience, Contact).

Build AppModel which manages:

current screen

theme

child models

✅ 4. Create Screen Navigation System

Each screen (Intro, Menu, Projects, Skills, etc.) is a separate Bubble Tea model.

Root model handles switching between screens.

Pressing ESC returns to Menu.

❗ Goal of Phase 2

You can switch between screens, even if they are empty.

PHASE 3 — Intro Splash: Typewriter + ASCII Portrait
✅ 5. Generate ASCII art

Options:

Use an ASCII generator locally.

Or generate via asciify-go on runtime.

Save result in assets/ascii.txt.

✅ 6. Build Typewriter Effect

Use Bubble Tea tickers.

Reveal intro text letter by letter.

Fade in ASCII portrait.

❗ Goal of Phase 3

You have a cinematic intro that leads into the menu when ENTER is pressed.

PHASE 4 — Command Palette (Your Navigation Hub)
✅ 7. Build Command Palette inspired by VS Code

Use bubbles/list with custom styling.

On pressing / anywhere → open palette.

Palette shows:

Projects

Experience

Skills

Contact me

Theme

❗ Goal of Phase 4

Your app is now navigable like a slick terminal app.

PHASE 5 — GitHub-Connected Projects
✅ 8. Create GitHub service

Use go-github or HTTP.

Fetch pinned repos or repos with specific topics.

✅ 9. Build Projects Screen

Show:

repo name

description

stars

language

Use bubbles/list for scrolling.

Press ENTER → show repo details with markdown rendered via Glamour.

❗ Goal of Phase 5

Your Projects section updates automatically from GitHub — huge flex.

PHASE 6 — Skills, Experience, Contact Screens
🛠 10. Build Skills Screen

Use short lists + Lipgloss styling.

Maybe add animation or icons.

🛠 11. Build Experience Screen

Use go-pretty/table or a vertical timeline style.

Show job titles, education, certifications.

🛠 12. Build Contact Screen

Display email, GitHub, LinkedIn.

Add a QR code (optional).

Add “press c to copy email” to clipboard (if env allows).

❗ Goal of Phase 6

Your content screens are visually consistent and clean.

PHASE 7 — Themes + Visual Polish
🎨 13. Build Theme Manager

Themes you can add:

Hacker Green

Solarized Dark

Monochrome Terminal

Dracula-ish

🔧 14. Add Theme Switcher

Press t → cycle themes.

OR choose theme inside the command palette.

🎭 15. Add Global Styling System

Define a component style system in internal/styles/:

titles

labels

borders

list items

background layouts

❗ Goal of Phase 7

Your TUI looks premium and consistent across all screens.

PHASE 8 — Interactive Easter Eggs (Optional But Cool)
🔥 16. Add Matrix Rain Animation

Use charmbracelet exp or implement your own.

Trigger with hidden command like matrix.

💻 17. Add Fake “Hacking Progress”

When user types hack:

Show a progress bar (bubbles/progress).

Fake logs appear on screen.

💬 18. Add Simple Chat About You

Ask questions:

“Who are you?”

“What tech do you use?”

“What’s your latest project?”

Answers appear in typewriter style.

❗ Goal of Phase 8

Your portfolio becomes memorable and fun.

PHASE 9 — SSH Deployment Setup
🧷 19. Build a small SSH server in Go

Users who SSH in activate the Bubble Tea TUI.

Wrap your Bubble Tea program in an SSH session.

🛠 20. Deploy on VPS

Set up a Droplet / Fly.io instance.

Configure SSH banner to trigger your TUI.

🚦 21. Secure your SSH entry

Disable shell access.

ForceCommand your-program.

Use separate user with restricted permissions.

❗ Goal of Phase 9

Anyone can run:

ssh jepoy@ssh.yourdomain.com


And boom — your portfolio opens. Biggest flex ever.

PHASE 10 — Final Polish
✨ 22. Test Across Terminals

Windows

macOS

Linux

Resize behavior

📘 23. Add README with instructions

How to SSH into your portfolio

Screenshot previews

Themes

Features

🧪 24. Add CI/CD

GitHub workflow that builds binaries

Deployment pipeline to your VPS

🎁 25. Final UX pass

Make animations smooth.

Improve spacing.

Remove flicker.

Tighten copywriting.