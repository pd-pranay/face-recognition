import { NbMenuItem } from '@nebular/theme';

export const MENU_ITEMS: NbMenuItem[] = [
  // {
  //   title: 'E-commerce',
  //   icon: 'shopping-cart-outline',
  //   link: '/pages/dashboard',
  //   home: true,
  // },
  // {
  //   title: 'IoT Dashboard',
  //   icon: 'home-outline',
  //   link: '/pages/iot-dashboard',
  // },
  {
    title: 'FEATURES',
    group: true,
  },
  {
    title: 'Find User',
    link: '/pages/forms/find-user',
  },
  {
    title: 'Add User',
    link: '/pages/forms/add-user',
  },
  {
    title: 'List User',
    link: '/pages/forms/list-user',
  },
  // {
  //   title: 'Face Recognition',
  //   icon: 'layout-outline',
  //   children: [
  //     {
  //       title: 'Find User',
  //       link: '/pages/forms/find-user',
  //     },
  //     {
  //       title: 'Add User',
  //       link: '/pages/forms/add-user',
  //     },
  //     {
  //       title: 'List User',
  //       link: '/pages/forms/list-user',
  //     },
  //   ],
  // },
];
