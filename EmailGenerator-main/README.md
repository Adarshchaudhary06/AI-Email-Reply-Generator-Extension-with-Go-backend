# AI Email Reply Extension 🚀

An AI-powered Chrome extension built with a custom **Golang Backend** and the **Groq API (Llama 3)** to generate instant, smart email replies directly inside Gmail.

## 📘 Description 
AI-powered Chrome extension using a custom lightweight Golang backend and Groq API to generate intelligent email replies with one-click convenience.

## 📂 Features
- Automatically adds an **AI Reply** button in Gmail's reply tab.  
- **Tone Selection:** Choose from Professional, Casual, Friendly, Direct, Concise, or Persuasive styles.
- Generates context-aware email responses instantly.  
- Seamless integration with Gmail UI.  
- Backend built with **Golang**.  
- Uses **Groq API** for natural language processing.
- Custom **Llama 3.3 70B** integration for high-quality responses.

## 🛠️ Tech Stack
- **Frontend:** JavaScript, Chrome Extensions API (Custom built UI injection)
- **Backend:** Golang  
- **AI Integration:** Groq API (Llama 3.3 70B)

## 🚀 How It Works
1. **Install the Extension:** Add the Chrome extension to your browser.  
2. **Open Gmail:** Go to your Gmail inbox.  
3. **Reply to an Email:** Click the reply button for any email.  
4. **Click 'AI Reply':** An additional button appears — click it to generate a reply.

## 🛠️ Installation
```bash
# Clone the repository
git clone https://github.com/selenophile1805/EmailGenerator.git

# Navigate to the go-backend
cd EmailGenerator-main/go-backend

# Run the Go application
go run .

# Load the extension into Chrome
1. Open Chrome and go to chrome://extensions/
2. Enable Developer Mode (toggle on top right)
3. Click "Load unpacked" and select the extension folder
```

## 🧩 API Configuration
1. Obtain Groq API credentials.  
2. You can configure the API Key in one of two ways:
   - **Option A (Recommended):** Create a file named `.env` inside the `go-backend` folder and add: `GROQ_API_KEY=your_actual_api_key_here`
   - **Option B (Terminal session):** Set the environment variable directly.
     - Example for Windows: `set GROQ_API_KEY=your_api_key`
     - Example for Mac/Linux: `export GROQ_API_KEY="your_api_key"`

## 🛠️ Usage
1. Open an email in Gmail.  
2. Click **Reply**.  
3. Hit the **AI Reply** button — your AI-generated response will appear!



