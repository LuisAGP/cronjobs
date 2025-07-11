document.addEventListener('DOMContentLoaded', function () {
    // Toggle between interval and cron options
    document.querySelectorAll('input[name="scheduleType"]').forEach(radio => {
        radio.addEventListener('change', function () {
            document.getElementById('intervalOptions').style.display =
                this.value === 'interval' ? 'block' : 'none';
            document.getElementById('cronOptions').style.display =
                this.value === 'cron' ? 'block' : 'none';
        });
    });

    // Show/hide request body based on HTTP method
    document.getElementById('httpMethod').addEventListener('change', function () {
        const showBody = ['POST', 'PUT', 'PATCH'].includes(this.value);
        document.getElementById('requestBodySection').style.display =
            showBody ? 'block' : 'none';
    });

    // Add/remove headers dynamically
    let headerCount = 1;
    document.getElementById('addHeader').addEventListener('click', function () {
        const container = document.getElementById('headersContainer');
        const newRow = document.createElement('div');
        newRow.className = 'header-row row mb-2';
        newRow.innerHTML = `<div class="col-md-5">
            <input type="text" class="form-control" name="headers[${headerCount}][key]" placeholder="Header name">
        </div>
        <div class="col-md-5">
            <input type="text" class="form-control" name="headers[${headerCount}][value]" placeholder="Header value">
        </div>
        <div class="col-md-2">
            <button type="button" class="btn btn-sm btn-outline-danger h-100" onclick="removeHeader(this)">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                    stroke-width="1.5" stroke="currentColor" width="20" height="20">
                    <path stroke-linecap="round" stroke-linejoin="round"
                        d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
                </svg>
            </button>
        </div>`;

        container.appendChild(newRow);
        headerCount++;
    });

    // Form submission
    document.getElementById('taskForm').addEventListener('submit', async function (e) {
        e.preventDefault();

        const formData = new FormData(this);
        const data = {};
        formData.forEach((value, key) => {
            data[key] = value;
        });

        try {
            const response = await fetch('/tasks', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            });

            if (response.ok) {
                // Close modal and refresh task list
                const modal = bootstrap.Modal.getInstance(document.getElementById('createTaskModal'));
                modal.hide();
                location.reload(); // Or update UI dynamically
            } else {
                alert('Error creating task');
            }
        } catch (error) {
            console.error('Error:', error);
            alert('Failed to create task');
        }
    });
});

function removeHeader(btn) {
    btn.closest('.header-row').remove();
}