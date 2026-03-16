console.log("AI Email Reply Extension specific content script loaded!");

// Configuration
const API_URL = "http://localhost:8080/email/generate";
const TONES = ["Professional", "Casual", "Friendly", "Direct", "Concise", "Persuasive"];

/**
 * Creates the UI elements (dropdown and button) to inject into the Gmail toolbar
 */
function createAIUI() {
    const container = document.createElement("div");
    container.className = "ai-reply-container";
    container.id = "ai-reply-container";

    // Create Tone Selector
    const toneSelect = document.createElement("select");
    toneSelect.className = "ai-tone-select";
    toneSelect.id = "ai-tone-select";

    TONES.forEach(tone => {
        const option = document.createElement("option");
        option.value = tone.toLowerCase();
        option.textContent = tone;
        toneSelect.appendChild(option);
    });

    // Create Generate Button
    const generateBtn = document.createElement("button");
    generateBtn.className = "ai-reply-button";
    generateBtn.innerHTML = "✨ AI Reply";
    generateBtn.type = "button";

    generateBtn.addEventListener("click", async () => {
        try {
            generateBtn.innerHTML = "⏳ Generating...";
            generateBtn.classList.add("loading");
            generateBtn.disabled = true;

            await handleGenerateReply(toneSelect.value);

        } catch (error) {
            console.error("AI Reply Error:", error);
            alert("Failed to generate reply. Ensure your backend is running on port 8080.");
        } finally {
            generateBtn.innerHTML = "✨ AI Reply";
            generateBtn.classList.remove("loading");
            generateBtn.disabled = false;
        }
    });

    container.appendChild(toneSelect);
    container.appendChild(generateBtn);
    return container;
}

/**
 * Grabs the original email content, calls the backend, and inserts the reply.
 */
async function handleGenerateReply(tone) {
    // 1. Get the original email content
    // Gmail creates div elements with 'gmail_quote' or specific classes for previous emails in thread
    const emailContents = document.querySelectorAll('.gE, .a3s, .gs');
    let originalEmail = "";

    if (emailContents.length > 0) {
        // usually the last one in the tree is the active one we are replying to
        const lastEmail = emailContents[emailContents.length - 1];
        originalEmail = lastEmail.innerText || lastEmail.textContent;
    }

    if (!originalEmail.trim()) {
        alert("Couldn't find an original email to reply to!");
        return;
    }

    // 2. Call the backend API
    const response = await fetch(API_URL, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            subject: "",
            emailContent: originalEmail,
            tone: tone
        })
    });

    if (!response.ok) {
        throw new Error(`API Error: ${response.status}`);
    }

    const generatedReply = await response.text();

    // 3. Insert the text into the Gmail compose box
    const composeBox = document.querySelector('[role="textbox"][g_editable="true"]');
    if (composeBox) {
        // Keep focus, format as HTML, and insert
        composeBox.focus();
        document.execCommand('insertHTML', false, generatedReply.replace(/\n/g, '<br>'));
    } else {
        console.error("Could not find the Gmail compose text box.");
        alert("Could not find the compose box to insert text.");
    }
}

/**
 * Observes the DOM to see when a Gmail Compose or Reply window is opened,
 * and injects our toolbar next to the Send button.
 */
const observer = new MutationObserver((mutations) => {
    for (const mutation of mutations) {
        if (mutation.addedNodes.length > 0) {
            // Looking for the Gmail toolbar row containing the Send button
            const toolbars = document.querySelectorAll('.btC, .aDh');

            toolbars.forEach(toolbar => {
                // To avoid duplicate injections check if we already added our container
                if (!toolbar.querySelector('#ai-reply-container')) {
                    // Inject our custom UI into this toolbar
                    toolbar.appendChild(createAIUI());
                }
            });
        }
    }
});

// Start observing the whole Gmail body for appended nodes
observer.observe(document.body, { childList: true, subtree: true });
