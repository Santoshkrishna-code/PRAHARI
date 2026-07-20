import React from 'react';
import { useForm, FieldValues } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { Button } from '../components/Button.tsx';
import { Input } from '../components/Input.tsx';

interface Field {
  name: string;
  label: string;
  type?: 'text' | 'password' | 'number';
  placeholder?: string;
}

interface DynamicFormProps {
  schema: z.ZodObject<any>;
  fields: Field[];
  onSubmit: (data: FieldValues) => void;
  submitLabel?: string;
}

export const DynamicForm: React.FC<DynamicFormProps> = ({
  schema,
  fields,
  onSubmit,
  submitLabel = 'Submit'
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting }
  } = useForm({
    resolver: zodResolver(schema)
  });

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4 w-full">
      {fields.map((field) => (
        <Input
          key={field.name}
          label={field.label}
          type={field.type}
          placeholder={field.placeholder}
          error={errors[field.name]?.message as string}
          {...register(field.name)}
        />
      ))}
      <Button type="submit" isLoading={isSubmitting} className="mt-2">
        {submitLabel}
      </Button>
    </form>
  );
};
