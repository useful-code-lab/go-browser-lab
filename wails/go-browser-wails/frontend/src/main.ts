import { Navigate, Reload, Home } from "../bindings/changeme/browserservice";

const address = document.querySelector<HTMLInputElement>("#address");
const addressForm = document.querySelector<HTMLFormElement>("#address-form");

const search = document.querySelector<HTMLInputElement>("#search");
const searchForm = document.querySelector<HTMLFormElement>("#search-form");

const startPage = document.querySelector<HTMLElement>("#start-page");

const backButton =
    document.querySelector<HTMLButtonElement>("#back");

const forwardButton =
    document.querySelector<HTMLButtonElement>("#forward");

const reloadButton =
    document.querySelector<HTMLButtonElement>("#reload");

const homeButton =
    document.querySelector<HTMLButtonElement>("#home");

const loading =
    document.querySelector<HTMLElement>("#loading");

function showLoading(): void {
    loading?.classList.add("active");
}

function hideLoading(): void {
    loading?.classList.remove("active");
}

async function navigate(value: string): Promise<void> {
    const query = value.trim();

    if (!query) {
        return;
    }

    showLoading();

    try {
        await Navigate(query);

        if (address) {
            address.value = query;
        }

        if (startPage) {
            startPage.style.display = "none";
        }
    } catch (error) {
        console.error("Navigation error:", error);
    } finally {
        window.setTimeout(hideLoading, 250);
    }
}

addressForm?.addEventListener("submit", async (event) => {
    event.preventDefault();

    if (!address) {
        return;
    }

    await navigate(address.value);
});

searchForm?.addEventListener("submit", async (event) => {
    event.preventDefault();

    if (!search) {
        return;
    }

    await navigate(search.value);
});

reloadButton?.addEventListener("click", async () => {
    showLoading();

    try {
        await Reload();
    } finally {
        window.setTimeout(hideLoading, 300);
    }
});

homeButton?.addEventListener("click", async () => {
    await Home();
});

backButton?.addEventListener("click", () => {
    console.log("Назад — добавим через browser history API");
});

forwardButton?.addEventListener("click", () => {
    console.log("Вперёд — добавим через browser history API");
});

/*
 * Быстрые ссылки.
 */

document
    .querySelectorAll<HTMLButtonElement>("[data-url]")
    .forEach((button) => {

        button.addEventListener("click", async () => {

            const value = button.dataset.url;

            if (value) {
                await navigate(value);
            }
        });
    });

/*
 * Ctrl+L — выделить адрес.
 * Ctrl+R — обновить.
 * Escape — убрать фокус.
 */

window.addEventListener("keydown", async (event) => {

    if (
        event.ctrlKey &&
        event.key.toLowerCase() === "l"
    ) {
        event.preventDefault();

        address?.focus();
        address?.select();

        return;
    }

    if (
        event.ctrlKey &&
        event.key.toLowerCase() === "r"
    ) {
        event.preventDefault();

        showLoading();

        try {
            await Reload();
        } finally {
            window.setTimeout(hideLoading, 300);
        }

        return;
    }

    if (event.key === "Escape") {
        address?.blur();
    }
});