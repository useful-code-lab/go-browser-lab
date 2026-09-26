const address = document.getElementById("address") as HTMLInputElement;
const goButton = document.getElementById("go") as HTMLButtonElement;

function normalizeUrl(value: string): string {
    const url = value.trim();

    if (!url) {
        return "https://example.com";
    }

    if (
        url.startsWith("http://") ||
        url.startsWith("https://")
    ) {
        return url;
    }

    return `https://${url}`;
}

function navigate(): void {
    const url = normalizeUrl(address.value);

    address.value = url;

    console.log("Navigate:", url);
}

goButton.addEventListener("click", navigate);

address.addEventListener("keydown", (event) => {
    if (event.key === "Enter") {
        navigate();
    }
});