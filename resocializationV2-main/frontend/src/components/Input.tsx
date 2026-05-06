import { forwardRef } from 'react';
import type { InputHTMLAttributes } from 'react';

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, className = '', ...props }, ref) => {
    return (
      <div className="form-group">
        {label && (
          <label className="form-label">
            {label}
          </label>
        )}
        <input
          ref={ref}
          className={`form-input ${error ? 'error' : ''} ${className}`.trim()}
          {...props}
        />
        {error && (
          <span className="form-error">{error}</span>
        )}
      </div>
    );
  }
);

Input.displayName = 'Input';

export default Input;
