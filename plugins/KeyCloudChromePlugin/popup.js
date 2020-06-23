function openDashboard() {
    chrome.tabs.executeScript({
        file: 'dashboard.js'
    });
}

document.getElementById('btnDashboard').addEventListener('click', openDashboard);
