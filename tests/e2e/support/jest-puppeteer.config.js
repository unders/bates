module.exports = {
    launch: {
        headless: true,
        // https://peter.sh/experiments/chromium-command-line-switches/
        args: ['--no-sandbox', '--disable-setuid-sandbox'],
    },
};
