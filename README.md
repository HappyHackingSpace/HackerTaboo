
# Hacker Taboo

Welcome to **Hacker Taboo!** This repository is dedicated to creating a fun and educational **Taboo-style card game** for hackers! Our goal is to compile a vast collection of cards covering topics like **computer science, cybersecurity, electronics, hacking, open source, and more!**

## 🚀 How to Play

1. Each card has a **keyword** that one player must describe.
2. The card also includes a **list of forbidden words** ("taboo words") that they **cannot** use while describing the keyword.
3. Their teammates must guess the keyword based on the description.
4. The team that guesses the most words correctly wins!

---

## 🤝 How to Contribute

We welcome contributions from everyone! Follow these steps to contribute new cards or improve existing ones.

### 1️⃣ Fork & Clone the Repository

1. Click the **Fork** button (top right of the repo) to create your own copy.
2. Clone your forked repository:
   ```sh
   git clone https://github.com/happyhackingspace/HackerTaboo.git
   cd HackerTaboo
   ```

### 2️⃣ Add a New Card

- Navigate to the `cards/` directory.
- Inside `cards/`, you will find category-based Markdown files:
  - `computer-science.md`
  - `cybersecurity.md`
  - `electronics.md`
  - `open-source.md`
- Open the relevant category file and add a new entry at the end using the following format:
  
  ```md
  ## Keyword: [Main term]
  
  **Taboo Words:** word1, word2, word3, word4, word5
  
  **Description:**
  A short and clear description of the keyword without using taboo words.
  ```

Example (inside `cybersecurity.md`):

```md
## Keyword: Firewall

**Taboo Words:** security, network, block, filter, protection

**Description:**
A system that monitors and controls incoming and outgoing network traffic based on predefined security rules.
```

### 3️⃣ Commit & Push Changes

```sh
git add cards/

git commit -m "Added new taboo card for [keyword] in [category]"

git push origin main
```

### 4️⃣ Create a Pull Request (PR)

1. Go to your forked repository on GitHub.
2. Click on **Pull Requests** > **New Pull Request**.
3. Select **compare across forks** and choose your fork.
4. Add a description of your contribution and submit the PR!

---

## 🛠 Contribution Guidelines

✅ Make sure your cards are **clear and concise**.
✅ Avoid using **taboo words** in the description.
✅ Maintain a **neutral and educational** tone.
✅ Ensure the keyword is relevant.
✅ Follow the existing format for consistency.

---

## 📜 License

This project is licensed under the [MIT License](LICENSE). Feel free to contribute and share!

---

## 🌎 Join the Community

💬 **Discussions:** [GitHub Discussions](https://github.com/YOUR-REPO/discussions)
 **Follow us on :**
- 🌐 [Website](https://happyhacking.space)
- 🎥 [YouTube](https://www.youtube.com/@HappyHackingSpace)
- 📷 [Instagram](https://www.instagram.com/happyhackingspace)
- 💼 [LinkedIn](https://www.linkedin.com/company/happyhackingspace)
- 🐦 [X (Twitter)](https://x.com/happyhackings)
- 🟦 [Bluesky](https://bsky.app/profile/happyhackingspace.bsky.social)
- 🐘 [Mastodon](https://mastodon.social/@happyhackingspace)
- 🔴 [Reddit](https://www.reddit.com/r/HappyHackingSpace/hot/)
- 🎮 [Twitch](https://www.twitch.tv/happyhackingspace)
- 🎵 [TikTok](https://www.tiktok.com/@happyhackingspace)
- 🟢 [Kick](https://kick.com/happyhackingspace)
- 👥 [Kommunity](https://kommunity.com/diyarbakir-happy-hacking-space)

Happy Hacking!