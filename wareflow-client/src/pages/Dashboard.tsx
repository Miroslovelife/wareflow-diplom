// src/components/Dashboard.tsx
import { useRole } from '../hooks/useRole';

export const Dashboard: React.FC = () => {
    const { isAdmin, isModerator } = useRole();

    return (
        <div className="dashboard">
            <h1>Панель управления</h1>

            {isAdmin() && (
                <section className="admin-section">
                    <h2>Административные функции</h2>
                    {/* Админ контент */}
                </section>
            )}

            {(isModerator() || isAdmin()) && (
                <section className="moderator-section">
                    <h2>Функции модератора</h2>
                    {/* Модератор контент */}
                </section>
            )}

            <section className="user-section">
                <h2>Общие функции</h2>
                {/* Общий контент */}
            </section>
        </div>
    );
};