package utils;

public class FlagUtils {
    private FlagUtils() {}

    public static String[] getFlags() {
        return new String[] {"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "--disable-dev-shm-usage",
            "--disable-features=site-per-process,TranslateUI,BlinkGenPropertyTrees", "--disable-hang-monitor",
            "--disable-ipc-flooding-protection", "--password-store=basic", "--use-mock-keychain",
            "--disable-backgrounding-occluded-windows", "--disable-breakpad",
            "--disable-client-side-phishing-detection", "--disable-prompt-on-repost", "--mute-audio",
            "--enable-features=NetworkService,NetworkServiceInProcess", "--hide-scrollbars", "--no-first-run",
            "--force-color-profile=srgb", "--disable-extensions", "--disable-background-timer-throttling",
            "--disable-popup-blocking", "--disable-sync", "--no-default-browser-check",
            "--disable-background-networking", "--metrics-recording-only", "--safebrowsing-disable-auto-update",
            "--disable-renderer-backgrounding", "--enable-automation", "--disable-default-apps",
            "--user-data-dir=/var/folders/cm/l1vj1p614fz7wjcdyx0604700000gp/T/chromedp-runner460319425",
            "--remote-debugging-port=0", "about:blank",};
    }
}
