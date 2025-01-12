import { ReactNode, useEffect, useState } from 'react';
import AnimateHide from '@kloudlite/design-system/atoms/animate-hide';
import { Button, IconButton } from '@kloudlite/design-system/atoms/button';
import { TextInput } from '@kloudlite/design-system/atoms/input';
import { cn, uuid } from '@kloudlite/design-system/utils';
import { MinusCircle, Plus } from '~/iotconsole/components/icons';

interface IKeyValuePair {
  onChange?(
    itemArray: Array<Record<string, any>>,
    itemObject: Record<string, any>
  ): void;
  value?: Array<Record<string, any>>;
  label?: ReactNode;
  message?: ReactNode;
  error?: boolean;
  size?: 'lg' | 'md';
  addText?: string;
  keyLabel?: string;
  valueLabel?: string;
}
const KeyValuePair = ({
  onChange,
  value = [],
  label,
  message,
  error,
  size,
  addText,
  keyLabel = 'key',
  valueLabel = 'value',
}: IKeyValuePair) => {
  const newItem = [{ [keyLabel]: '', [valueLabel]: '', id: uuid() }];
  const [items, setItems] = useState<Array<Record<string, any>>>(newItem);

  const handleChange = (_value = '', id = '', target = {}) => {
    const tempItems = items.map((i) => {
      if (i.id === id) {
        switch (target) {
          case 'key':
            return { ...i, [keyLabel]: _value };

          case 'value':
          default:
            return { ...i, [valueLabel]: _value };
        }
      }
      return i;
    });
    const formatItems = tempItems.reduce((acc, curr) => {
      if (curr.key && curr.value) {
        acc[curr.key] = curr.value;
      }
      return acc;
    }, {});
    if (onChange) onChange(Array.from(tempItems), formatItems);
  };

  useEffect(() => {
    if (value && value.length === 0) {
      setItems(newItem);
      return;
    }
    setItems(
      Array.from(value || newItem).map((v) => ({
        ...v,
        id: v.id ? v.id : uuid(),
      }))
    );
  }, [value]);

  return (
    <div className="flex flex-col">
      <div className="flex flex-col">
        <div className="flex flex-col gap-md">
          {label && (
            <span className="text-text-default bodyMd-medium">{label}</span>
          )}
          {items.map((item) => (
            <div key={item.id} className="flex flex-row gap-xl items-start">
              <div className="flex-1">
                <TextInput
                  size={size || 'md'}
                  error={error}
                  placeholder="Key"
                  value={item[keyLabel]}
                  onChange={({ target }) =>
                    handleChange(target.value, item.id, 'key')
                  }
                />
              </div>
              <div className="flex-1">
                <TextInput
                  size={size || 'md'}
                  error={error}
                  placeholder="Value"
                  value={item[valueLabel]}
                  onChange={({ target }) =>
                    handleChange(target.value, item.id, 'value')
                  }
                />
              </div>
              <div className="self-center">
                <IconButton
                  icon={<MinusCircle />}
                  variant="plain"
                  disabled={items.length < 2}
                  onClick={() => {
                    setItems(items.filter((i) => i.id !== item.id));
                  }}
                />
              </div>
            </div>
          ))}
        </div>
        <AnimateHide show={!!message}>
          <div
            className={cn(
              'bodySm pulsable',
              {
                'text-text-critical': !!error,
                'text-text-default': !error,
              },
              'pt-md'
            )}
          >
            {message}
          </div>
        </AnimateHide>
        <div className="pt-xl">
          <Button
            variant="basic"
            content={addText || 'Add'}
            size="sm"
            prefix={<Plus />}
            onClick={() => {
              setItems([...items, { ...newItem[0], id: uuid() }]);
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default KeyValuePair;
