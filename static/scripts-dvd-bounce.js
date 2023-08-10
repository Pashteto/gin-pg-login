document.addEventListener('DOMContentLoaded', (event) => {

    let sandwich1 = document.getElementById('sandwich1');
    let sandwich2 = document.getElementById('sandwich2');

    let deltaX1 = 2;
    let deltaY1 = 2;
    let deltaX2 = -2;
    let deltaY2 = -2;

    // Initial positions for sandwiches
    let x1 = Math.random() * (window.innerWidth - sandwich1.offsetWidth);
    let y1 = Math.random() * (window.innerHeight - sandwich1.offsetHeight);
    let x2 = Math.random() * (window.innerWidth - sandwich2.offsetWidth);
    let y2 = Math.random() * (window.innerHeight - sandwich2.offsetHeight);

    sandwich1.style.left = x1 + 'px';
    sandwich1.style.top = y1 + 'px';
    sandwich2.style.left = x2 + 'px';
    sandwich2.style.top = y2 + 'px';

    function moveSandwich() {
        // Check bounds for sandwich1
        if (x1 + sandwich1.offsetWidth > window.innerWidth || x1 < 0) deltaX1 = -deltaX1;
        if (y1 + sandwich1.offsetHeight > window.innerHeight || y1 < 0) deltaY1 = -deltaY1;

        // Check bounds for sandwich2
        if (x2 + sandwich2.offsetWidth > window.innerWidth || x2 < 0) deltaX2 = -deltaX2;
        if (y2 + sandwich2.offsetHeight > window.innerHeight || y2 < 0) deltaY2 = -deltaY2;

        x1 += deltaX1;
        y1 += deltaY1;
        x2 += deltaX2;
        y2 += deltaY2;

        sandwich1.style.left = x1 + 'px';
        sandwich1.style.top = y1 + 'px';
        sandwich2.style.left = x2 + 'px';
        sandwich2.style.top = y2 + 'px';

        requestAnimationFrame(moveSandwich);
    }

    moveSandwich();
});
