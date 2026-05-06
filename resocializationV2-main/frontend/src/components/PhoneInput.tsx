import { forwardRef } from 'react';
import type { InputHTMLAttributes } from 'react';
import { formatPhone } from '../utils/phone';

interface PhoneInputProps extends Omit<InputHTMLAttributes<HTMLInputElement>, 'onChange'> {
  label?: string;
  error?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const PhoneInput = forwardRef<HTMLInputElement, PhoneInputProps>(
  ({ label, error, className = '', onChange, value, ...props }, ref) => {
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      const formatted = formatPhone(e.target.value);

      // Cria um novo evento com o valor formatado
      const newEvent = {
        ...e,
        target: {
          ...e.target,
          value: formatted,
        },
      } as React.ChangeEvent<HTMLInputElement>;

      if (onChange) {
        onChange(newEvent);
      }
    };

    return (
      <div className="form-group">
        {label && (
          <label className="form-label">
            {label}
          </label>
        )}
        <input
          ref={ref}
          type="tel"
          className={`form-input ${error ? 'error' : ''} ${className}`.trim()}
          value={value}
          onChange={handleChange}
          placeholder="(00) 00000-0000"
          maxLength={15}
          {...props}
        />
        {error && (
          <span className="form-error">{error}</span>
        )}
      </div>
    );
  }
);

PhoneInput.displayName = 'PhoneInput';

export default PhoneInput;
