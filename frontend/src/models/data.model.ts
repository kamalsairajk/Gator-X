export interface UserTooltipData {
  username: string,
  email: string,
}

export interface UserData {
  name: string,
  username: string,
  email: string,
  phone: string,
  id: any,
  website: string,
  address: string,
  company: string,
}

export interface UserId {
  userId: any,
}

export interface AlbumId {
  albumId: any;
}

export interface PostId {
  postId: any,
}

export interface NewPost {
  title: string,
  body: string,
  userId: number,
}

export interface Post {
  id: number,
  title: string,
  body: string,
  userId: number,
}

export interface Id {
  id: number,
}
