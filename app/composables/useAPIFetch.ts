export const useAPIFetch: typeof useFetch = (request, opts?) => {
  const config = useRuntimeConfig();

  return useFetch(request, {
    baseURL: config.public.baseURL as string,
    ...opts,
  });
};
