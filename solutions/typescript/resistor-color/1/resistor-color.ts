export const colorCode = (color?: ColorID): number | ColorID[] => {
  if (color === undefined) return COLORS;
  return colorId[color];
};

export type ColorID = "black" | "brown" | "red" | "orange" | "yellow" | "green" | "blue" | "violet" | "grey" | "white";

export const COLORS: ColorID[] = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"];

export const colorId: Record<ColorID, number> = {
  "black": 0,
  "brown": 1,
  "red": 2,
  "orange": 3,
  "yellow": 4,
  "green": 5,
  "blue": 6,
  "violet": 7,
  "grey": 8,
  "white": 9,
};