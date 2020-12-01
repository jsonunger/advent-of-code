export default () => {
  const today = new Date();
  const day = today.getDate();
  const year = today.getFullYear();
  return `${year}/day${day}`;
};
