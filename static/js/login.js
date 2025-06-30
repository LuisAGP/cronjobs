document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');
    
    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;
            
            try {
                const response = await fetch('/api/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ email, password })
                });
                
                const data = await response.json();
                
                if (response.ok) {
                    // Guardar tokens en localStorage
                    localStorage.setItem('accessToken', data.access_token);
                    localStorage.setItem('refreshToken', data.refresh_token);
                    
                    // Redirigir al dashboard
                    window.location.href = '/dashboard';
                } else {
                    showError(data.error || 'Error en la autenticación');
                }
            } catch (error) {
                showError('Error de conexión');
            }
        });
    }
    
    function showError(message) {
        // Eliminar mensajes anteriores
        const oldError = document.querySelector('.error-message');
        if (oldError) oldError.remove();
        
        const errorDiv = document.createElement('div');
        errorDiv.className = 'error-message';
        errorDiv.textContent = message;
        
        const form = document.getElementById('loginForm');
        if (form) {
            form.parentNode.insertBefore(errorDiv, form);
        }
    }
});