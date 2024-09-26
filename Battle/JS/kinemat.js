// script.js
const grid = document.getElementById('grid');

// Создаем клетки
for (let i = 0; i < 100; i++) {
    const cell = document.createElement('div');
    cell.addEventListener('dragover', (e) => {
        e.preventDefault();
    });
    cell.addEventListener('drop', (e) => {
        const id = e.dataTransfer.getData('text/plain');
        const draggedElement = document.getElementById(id);
        cell.appendChild(draggedElement);
    });
    grid.appendChild(cell);
}

// Устанавливаем обработчики для объектов
const objects = document.querySelectorAll('.object');
objects.forEach(obj => {
    obj.addEventListener('dragstart', (e) => {
        e.dataTransfer.setData('text/plain', obj.id);
    });
});
