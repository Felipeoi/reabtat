import type { ButtonHTMLAttributes, ReactNode } from 'react';

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'danger' | 'ghost';
  size?: 'sm' | 'md' | 'lg';
  children: ReactNode;
}

export default function Button({
  variant = 'primary',
  size = 'md',
  className = '',
  children,
  ...props
}: ButtonProps) {
  const variantClass = variant === 'primary' ? 'btn-primary' :
                      variant === 'secondary' ? 'btn-secondary' :
                      variant === 'danger' ? 'btn-danger' :
                      variant === 'ghost' ? 'btn-ghost' : '';

  const sizeClass = size === 'sm' ? 'btn-sm' :
                   size === 'lg' ? 'btn-lg' : '';

  return (
    <button
      className={`btn ${variantClass} ${sizeClass} ${className}`.trim()}
      {...props}
    >
      {children}
    </button>
  );
}
