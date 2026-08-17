import { Pressable } from "react-native";

interface ButtonProps {
  children: React.ReactNode;
  onPress: () => void;
  variant?: 'default' | 'outline';
  className?: string;
  disabled?: boolean;
}



export function Button({ children, onPress, variant = 'default', className = '', disabled }: ButtonProps) {

  const buttonVariants = {
    default: `bg-dark-blue-500 active:bg-dark-blue-600 dark:bg-primary-500 dark:active:bg-primary-600 active:opacity-75 font-bold text-white ${disabled ? 'opacity-50' : ''}`,
    outline: `border border-dark-blue-500 text-dark-blue-500 dark:text-primary-500 dark:font-bold active:bg-dark-blue-600/10 active:opacity-75 ${disabled ? 'opacity-50' : ''}`,
  };

  return (
    <Pressable
      className={`${buttonVariants[variant]} p-[1em] items-center rounded-md ${className}`}
      onPress={onPress}
      disabled={disabled}
    >
      {children}
    </Pressable>
  );
}