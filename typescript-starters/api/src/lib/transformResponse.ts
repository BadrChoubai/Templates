export function transformResponse<TIn, TOut>(
  data: TIn,
  transformFn: (data: TIn) => TOut,
): TOut {
  return transformFn(data);
}
