import { PluginFunction, Types } from '@graphql-codegen/plugin-helpers';
import { GraphQLSchema, isEnumType } from 'graphql';

export const plugin: PluginFunction = (
  schema: GraphQLSchema,
  documents: Types.DocumentFile[],
  config: any
) => {
  console.log('hllow');
  const allTypes = Object.values(schema.getTypeMap());
  const enumTypes = allTypes.filter(isEnumType);

  // const result = enumTypes
  //   .map((enumType) => {
  //     const values = enumType
  //       .getValues()
  //       .map((value) => `'${value.name}'`)
  //       .join(' | ');
  //     return `export type ${enumType.name} = ${values} | string;`;
  //   })
  //   .join('\n');

  return { prepend: [], content: '' };
};
