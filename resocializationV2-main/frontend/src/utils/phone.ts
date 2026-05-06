export const formatPhone = (value: string): string => {
  // Remove tudo que não é número
  const numbers = value.replace(/\D/g, '');

  // Limita a 11 dígitos
  const limited = numbers.slice(0, 11);

  // Formata baseado no tamanho
  if (limited.length <= 2) {
    return limited;
  } else if (limited.length <= 6) {
    // (XX) XXXX
    return `(${limited.slice(0, 2)}) ${limited.slice(2)}`;
  } else if (limited.length <= 10) {
    // (XX) XXXX-XXXX
    return `(${limited.slice(0, 2)}) ${limited.slice(2, 6)}-${limited.slice(6)}`;
  } else {
    // (XX) XXXXX-XXXX
    return `(${limited.slice(0, 2)}) ${limited.slice(2, 7)}-${limited.slice(7, 11)}`;
  }
};

export const unformatPhone = (value: string): string => {
  return value.replace(/\D/g, '');
};
