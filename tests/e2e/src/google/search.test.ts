describe('Google', () => {
    beforeAll(async () => {
        await page.goto('https://google.com');
    });


    test('should be on page: "google.com"', async () => {
        await expect(page).toMatch('google')
    }, 5000);

    test('search should find "unders" when searching for: "github Anders Törnqvist"', async () => {
        const el = await page.$('input[name="q"]');
        expect(el).not.toBeNull();
        if (el == null) {
            return;
        }
        await el.type("github Anders Törnqvist");
        await el.press("Enter");

        await page.waitForNavigation({ waitUntil: "networkidle2" });
        await page.waitFor("#bfoot");
        await expect(page).toMatch("unders");
    }, 10000);

});
