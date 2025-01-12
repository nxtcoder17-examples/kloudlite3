import { ReactNode } from 'react';

export type NonNullableString = string & NonNullable<undefined>;

export type MapType<T = string | number | Date | boolean> = {
  [key: string]: T | MapType<T>;
};

export type FlatMapType<T = string | number | boolean> = {
  [key: string]: T;
};

export interface IChildren {
  children: ReactNode;
}

export interface IRemixHeader {
  get?: any;
}

export interface IRemixReq {
  headers: IRemixHeader;
  url: string;
  method: 'GET' | 'POST' | (string & NonNullable<unknown>);
  json: () => Promise<MapType>;
  cookies: string[];
}

export interface IRemixCtx {
  request: IRemixReq;
  params: FlatMapType<string>;
}

export interface IExtRemixCtx extends IRemixCtx {
  authProps: any;
  consoleContextProps: any;
  accounts: any;
  request: IRemixReq;
}

export type ICookie = any;
export type ICookies = ICookie[];

export interface IGQLServerProps {
  headers: IRemixHeader;
  cookies?: ICookies;
}

type ROnly<T> = {
  readonly [k in keyof T]: T[k] extends object ? ROnly<T[k]> : T[k];
};

export type NN<T> = NonNullable<T>;

export type RNN<T> = {
  readonly [k in keyof T]: NN<T[k]>;
};

export type DeepReadOnly<T> = ROnly<T>;

export type IGqlReturn<T> = Promise<{
  errors?: Error[];
  data: T;
}>;

export type ExtractArrayType<T> = T extends (infer U)[] ? U : never;

export type LoaderResult<T extends (props: any) => Promise<any>> = Awaited<
  ReturnType<T>
>;

export type ISetState<T> = React.Dispatch<React.SetStateAction<T>>;
