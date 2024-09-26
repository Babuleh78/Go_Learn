

const EnemyGrid = document.getElementById('EnemyGrid');

const CreateEnemyGrid = () => {

    for (let i = 0; i < 100; i++) {
        const cell = document.createElement('div');
        cell.className = 'cell'; 
        cell.addEventListener('click', () => {
            cell.classList.toggle('active'); 
        });
        EnemyGrid.appendChild(cell); 
    }
};

CreateEnemyGrid(); 
