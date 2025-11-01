import { Chessboard } from 'react-chessboard';
import { Chess } from 'chess.js';
import { useState, useEffect } from 'react';
import { useMatchStore } from '@/stores/match-store';
import { sendPlayMove } from '@/events/play-move-event';

export const ChessBoard = () => {
  const { boardOrientation, position, playerColor, status } = useMatchStore();
  const [game, setGame] = useState(new Chess());

  useEffect(() => {
    if (position && position !== 'start') {
      const newGame = new Chess(position);
      setGame(newGame);
    } else {
      setGame(new Chess());
    }
  }, [position]);

  const onPieceDrop = ({ sourceSquare, targetSquare }: { piece: any; sourceSquare: string; targetSquare: string | null }) => {
    if (status !== 'playing' || !targetSquare) return false;

    const currentTurn = game.turn() === 'w' ? 'white' : 'black';
    if (currentTurn !== playerColor) {
      return false;
    }

    try {
      const move = game.move({
        from: sourceSquare,
        to: targetSquare,
        promotion: 'q',
      });

      if (move === null) return false;

      const newFen = game.fen();
      setGame(new Chess(newFen));
      useMatchStore.getState().updatePosition(newFen);

      sendPlayMove({ from: sourceSquare, to: targetSquare });

      return true;
    } catch (error) {
      return false;
    }
  };

  return (
    <div className="w-full max-w-[600px]">
      <Chessboard
        options={{
          position: game.fen(),
          onPieceDrop: onPieceDrop,
          boardOrientation: boardOrientation,
          boardStyle: {
            borderRadius: '4px',
            boxShadow: '0 2px 10px rgba(0, 0, 0, 0.5)',
          },
        }}
      />
    </div>
  );
}