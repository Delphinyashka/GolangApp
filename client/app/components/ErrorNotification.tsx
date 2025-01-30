"use client";
import { Notification } from "@mantine/core";
import { IconX } from "@tabler/icons-react";

interface ErrorNotificationProps {
    errorMessage: string | null;
    onClose: () => void;
}

const ErrorNotification: React.FC<ErrorNotificationProps> = ({ errorMessage, onClose }) => {
    if (!errorMessage) return null;

    return (
        <Notification
            color="red"
            icon={<IconX size="1.1rem" />}
            style={{
                position: "fixed",
                top: "20px",
                right: "20px",
                zIndex: 999,
            }}
            onClose={onClose}
        >
            {errorMessage}
        </Notification>
    );
};

export default ErrorNotification;
