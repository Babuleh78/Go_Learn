const socket = new WebSocket("ws://localhost:8080/ws");
    

const MyGrid = document.getElementById('MyGrid');
        const CreateMyGrid = () => {
            for (let i = 0; i < 100; i++) {
                const cell = document.createElement('div');
                cell.className = 'cell';
                cell.addEventListener('click', () => {
                    cell.classList.toggle('active');
                });
                MyGrid.appendChild(cell);
            }
        };

        CreateMyGrid();