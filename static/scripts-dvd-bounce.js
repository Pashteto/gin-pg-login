const sandwich1 = document.getElementById('sandwich1');
const sandwich2 = document.getElementById('sandwich2');

function moveElement(element) {
    let x = 0, y = 0;
    let xSpeed = 2, ySpeed = 2;

    function move() {
        if (x + element.clientWidth > window.innerWidth || x < 0) {
            xSpeed = -xSpeed;
        }

        if (y + element.clientHeight > window.innerHeight || y < 0) {
            ySpeed = -ySpeed;
        }

        x += xSpeed;
        y += ySpeed;

        element.style.left = x + 'px';
        element.style.top = y + 'px';

        requestAnimationFrame(move);
    }

    move();
}

moveElement(sandwich1);
moveElement(sandwich2);