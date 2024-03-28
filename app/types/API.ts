export type BaseAPIResponse = {
  success: boolean;
  message: string;
  data: any;
};

export type APIBookResponse = {
  volumeInfo: {
    authors: string[];
    title: string;
    description: string;
    imageLinks: {
      thumbnail: string;
    };
    categories: string[];
    industryIndentifiers: {
      type: string;
      identifier: string;
    };
    pageCount: number;
    publishedDate: string;
  };
};
